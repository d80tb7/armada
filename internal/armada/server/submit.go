package server

import (
	"context"
	"fmt"
	"github.com/G-Research/armada/internal/common"
	"github.com/G-Research/armada/internal/common/util"
	"github.com/G-Research/armada/internal/common/validation"
	"github.com/gogo/protobuf/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"strings"
	"time"

	"github.com/G-Research/armada/internal/armada/configuration"
	"github.com/G-Research/armada/internal/armada/permissions"
	"github.com/G-Research/armada/internal/armada/repository"
	"github.com/G-Research/armada/internal/common/auth/authorization"
	"github.com/G-Research/armada/internal/common/auth/permission"
	"github.com/G-Research/armada/pkg/api"
)

type SubmitServer struct {
	permissions              authorization.PermissionChecker
	jobRepository            repository.JobRepository
	queueRepository          repository.QueueRepository
	eventStore               repository.EventStore
	schedulingInfoRepository repository.SchedulingInfoRepository
	queueManagementConfig    *configuration.QueueManagementConfig
	defaultJobLimits         common.ComputeResources
	defaultJobTolerations    []v1.Toleration
}

func NewSubmitServer(
	permissions authorization.PermissionChecker,
	jobRepository repository.JobRepository,
	queueRepository repository.QueueRepository,
	eventStore repository.EventStore,
	schedulingInfoRepository repository.SchedulingInfoRepository,
	queueManagementConfig *configuration.QueueManagementConfig,
	defaultJobLimits common.ComputeResources,
	defaultJobTolerations []v1.Toleration) *SubmitServer {

	return &SubmitServer{
		permissions:              permissions,
		jobRepository:            jobRepository,
		queueRepository:          queueRepository,
		eventStore:               eventStore,
		schedulingInfoRepository: schedulingInfoRepository,
		queueManagementConfig:    queueManagementConfig,
		defaultJobLimits:         defaultJobLimits,
		defaultJobTolerations:    defaultJobTolerations}
}

func (server *SubmitServer) GetQueueInfo(ctx context.Context, req *api.QueueInfoRequest) (*api.QueueInfo, error) {
	if e := checkPermission(server.permissions, ctx, permissions.WatchAllEvents); e != nil {
		return nil, e
	}
	jobSets, e := server.jobRepository.GetQueueActiveJobSets(req.Name)
	if e != nil {
		return nil, e
	}
	return &api.QueueInfo{
		Name:          req.Name,
		ActiveJobSets: jobSets,
	}, nil
}

func (server *SubmitServer) GetQueue(ctx context.Context, req *api.QueueGetRequest) (*api.Queue, error) {
	queue, e := server.queueRepository.GetQueue(req.Name)
	if e == repository.ErrQueueNotFound {
		return nil, status.Errorf(codes.NotFound, "Queue %q not found", req.Name)

	} else if e != nil {
		return nil, status.Errorf(codes.Unavailable, "Could not load queue %q: %s", req.Name, e.Error())
	}
	return queue, nil
}

func (server *SubmitServer) CreateQueue(ctx context.Context, queue *api.Queue) (*types.Empty, error) {
	if e := checkPermission(server.permissions, ctx, permissions.CreateQueue); e != nil {
		return nil, e
	}

	if len(queue.UserOwners) == 0 {
		principal := authorization.GetPrincipal(ctx)
		queue.UserOwners = []string{principal.GetName()}
	}

	e := validateQueue(queue)
	if e != nil {
		return nil, e
	}

	e = server.queueRepository.CreateQueue(queue)
	if e == repository.ErrQueueAlreadyExists {
		return nil, status.Errorf(codes.AlreadyExists, "Queue %q already exists", queue.Name)
	} else if e != nil {
		return nil, status.Errorf(codes.Unavailable, e.Error())
	}
	return &types.Empty{}, nil
}

func (server *SubmitServer) UpdateQueue(ctx context.Context, queue *api.Queue) (*types.Empty, error) {
	if e := checkPermission(server.permissions, ctx, permissions.CreateQueue); e != nil {
		return nil, e
	}

	e := validateQueue(queue)
	if e != nil {
		return nil, e
	}

	e = server.queueRepository.UpdateQueue(queue)
	if e == repository.ErrQueueNotFound {
		return nil, status.Errorf(codes.NotFound, "Queue %q not found", queue.Name)
	} else if e != nil {
		return nil, status.Errorf(codes.Unavailable, e.Error())
	}
	return &types.Empty{}, nil
}

func (server *SubmitServer) DeleteQueue(ctx context.Context, request *api.QueueDeleteRequest) (*types.Empty, error) {
	if e := checkPermission(server.permissions, ctx, permissions.DeleteQueue); e != nil {
		return nil, e
	}

	active, e := server.jobRepository.GetQueueActiveJobSets(request.Name)
	if e != nil {
		return nil, status.Errorf(codes.InvalidArgument, e.Error())
	}
	if len(active) > 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "Queue is not empty.")
	}

	e = server.queueRepository.DeleteQueue(request.Name)
	if e != nil {
		return nil, status.Errorf(codes.InvalidArgument, e.Error())
	}

	return &types.Empty{}, nil
}

func (server *SubmitServer) SubmitJobs(ctx context.Context, req *api.JobSubmitRequest) (*api.JobSubmitResponse, error) {
	e, ownershipGroups := server.checkQueuePermission(ctx, req.Queue, true, permissions.SubmitJobs, permissions.SubmitAnyJobs)
	if e != nil {
		return nil, e
	}

	principal := authorization.GetPrincipal(ctx)

	jobs, e := server.createJobs(req, principal.GetName(), ownershipGroups)
	if e != nil {
		return nil, status.Errorf(codes.InvalidArgument, e.Error())
	}

	allClusterSchedulingInfo, e := server.schedulingInfoRepository.GetClusterSchedulingInfo()
	if e != nil {
		return nil, e
	}

	e = validateJobsCanBeScheduled(jobs, allClusterSchedulingInfo)
	if e != nil {
		return nil, status.Errorf(codes.InvalidArgument, e.Error())
	}

	e = reportSubmitted(server.eventStore, jobs)
	if e != nil {
		return nil, status.Errorf(codes.Aborted, e.Error())
	}

	submissionResults, e := server.jobRepository.AddJobs(jobs)
	if e != nil {
		return nil, status.Errorf(codes.Aborted, e.Error())
	}

	result := &api.JobSubmitResponse{
		JobResponseItems: make([]*api.JobSubmitResponseItem, 0, len(submissionResults)),
	}

	createdJobs := []*api.Job{}
	doubleSubmits := []*repository.SubmitJobResult{}
	for i, submissionResult := range submissionResults {
		jobResponse := &api.JobSubmitResponseItem{JobId: submissionResult.JobId}
		if submissionResult.Error != nil {
			jobResponse.Error = submissionResult.Error.Error()
		}
		result.JobResponseItems = append(result.JobResponseItems, jobResponse)

		if submissionResult.Error == nil {
			if submissionResult.DuplicateDetected {
				doubleSubmits = append(doubleSubmits, submissionResult)
			} else {
				createdJobs = append(createdJobs, jobs[i])
			}
		}
	}

	e = reportDuplicateDetected(server.eventStore, doubleSubmits)
	if e != nil {
		return result, status.Errorf(codes.Internal, e.Error())
	}

	e = reportQueued(server.eventStore, createdJobs)
	if e != nil {
		return result, status.Errorf(codes.Internal, e.Error())
	}
	return result, nil
}

func (server *SubmitServer) CancelJobs(ctx context.Context, request *api.JobCancelRequest) (*api.CancellationResult, error) {
	if request.JobId != "" {
		jobs, e := server.jobRepository.GetExistingJobsByIds([]string{request.JobId})
		if e != nil {
			return nil, status.Errorf(codes.Internal, e.Error())
		}
		return server.cancelJobs(ctx, jobs[0].Queue, jobs)
	}

	if request.JobSetId != "" && request.Queue != "" {
		ids, e := server.jobRepository.GetActiveJobIds(request.Queue, request.JobSetId)
		if e != nil {
			return nil, status.Errorf(codes.Aborted, e.Error())
		}
		jobs, e := server.jobRepository.GetExistingJobsByIds(ids)
		if e != nil {
			return nil, status.Errorf(codes.Internal, e.Error())
		}
		return server.cancelJobs(ctx, request.Queue, jobs)
	}
	return nil, status.Errorf(codes.InvalidArgument, "Specify job id or queue with job set id")
}

func (server *SubmitServer) cancelJobs(ctx context.Context, queue string, jobs []*api.Job) (*api.CancellationResult, error) {
	if e, _ := server.checkQueuePermission(ctx, queue, false, permissions.CancelJobs, permissions.CancelAnyJobs); e != nil {
		return nil, e
	}
	principal := authorization.GetPrincipal(ctx)

	e := reportJobsCancelling(server.eventStore, principal.GetName(), jobs)
	if e != nil {
		return nil, status.Errorf(codes.Unknown, e.Error())
	}

	deletionResult := server.jobRepository.DeleteJobs(jobs)
	cancelled := []*api.Job{}
	cancelledIds := []string{}
	for job, err := range deletionResult {
		if err != nil {
			log.Errorf("Error when cancelling job id %s: %s", job.Id, err.Error())
		} else {
			cancelled = append(cancelled, job)
			cancelledIds = append(cancelledIds, job.Id)
		}
	}

	e = reportJobsCancelled(server.eventStore, principal.GetName(), cancelled)
	if e != nil {
		return nil, status.Errorf(codes.Unknown, e.Error())
	}

	return &api.CancellationResult{cancelledIds}, nil
}

// Returns mapping from job id to error (if present), for all existing jobs
func (server *SubmitServer) ReprioritizeJobs(ctx context.Context, request *api.JobReprioritizeRequest) (*api.JobReprioritizeResponse, error) {
	var jobs []*api.Job
	if len(request.JobIds) > 0 {
		existingJobs, err := server.jobRepository.GetExistingJobsByIds(request.JobIds)
		if err != nil {
			return nil, err
		}
		jobs = existingJobs
	} else if request.Queue != "" && request.JobSetId != "" {
		ids, e := server.jobRepository.GetActiveJobIds(request.Queue, request.JobSetId)
		if e != nil {
			return nil, status.Errorf(codes.Aborted, e.Error())
		}
		existingJobs, e := server.jobRepository.GetExistingJobsByIds(ids)
		if e != nil {
			return nil, status.Errorf(codes.Internal, e.Error())
		}
		jobs = existingJobs
	}

	err := server.checkReprioritizePerms(ctx, jobs)
	if err != nil {
		return nil, err
	}

	principalName := authorization.GetPrincipal(ctx).GetName()

	err = reportJobsReprioritizing(server.eventStore, principalName, jobs, request.NewPriority)
	if err != nil {
		return nil, err
	}

	jobIds := []string{}
	for _, job := range jobs {
		jobIds = append(jobIds, job.Id)
	}
	results, err := server.reprioritizeJobs(jobIds, request.NewPriority, principalName)
	if err != nil {
		return nil, err
	}

	return &api.JobReprioritizeResponse{ReprioritizationResults: results}, nil
}

func (server *SubmitServer) reprioritizeJobs(jobIds []string, newPriority float64, principalName string) (map[string]string, error) {
	updateJobResults := server.jobRepository.UpdateJobs(jobIds, func(jobs []*api.Job) {
		for _, job := range jobs {
			job.Priority = newPriority
		}
		err := server.reportReprioritizedJobEvents(jobs, newPriority, principalName)
		if err != nil {
			log.Warnf("Failed to report events for reprioritize of jobs %s: %v", strings.Join(jobIds, ", "), err)
		}
	})

	results := map[string]string{}
	for _, r := range updateJobResults {
		if r.Error == nil {
			results[r.JobId] = ""
		} else {
			results[r.JobId] = r.Error.Error()
		}
	}
	return results, nil
}

func (server *SubmitServer) reportReprioritizedJobEvents(reprioritizedJobs []*api.Job, newPriority float64, principalName string) error {

	err := reportJobsUpdated(server.eventStore, principalName, reprioritizedJobs)
	if err != nil {
		return err
	}

	err = reportJobsReprioritized(server.eventStore, principalName, reprioritizedJobs, newPriority)
	if err != nil {
		return err
	}

	return nil
}

func (server *SubmitServer) checkReprioritizePerms(ctx context.Context, jobs []*api.Job) error {
	queues := make(map[string]bool)
	for _, job := range jobs {
		queues[job.Queue] = true
	}
	for queue := range queues {
		if e, _ := server.checkQueuePermission(ctx, queue, false, permissions.ReprioritizeJobs, permissions.ReprioritizeAnyJobs); e != nil {
			return e
		}
	}
	return nil
}

func (server *SubmitServer) checkQueuePermission(
	ctx context.Context,
	queueName string,
	attemptToCreate bool,
	basicPermission permission.Permission,
	allQueuesPermission permission.Permission) (e error, ownershipGroups []string) {

	queue, e := server.queueRepository.GetQueue(queueName)
	if e == repository.ErrQueueNotFound {
		if attemptToCreate &&
			server.queueManagementConfig.AutoCreateQueues &&
			server.permissions.UserHasPermission(ctx, permissions.SubmitAnyJobs) {

			queue = &api.Queue{
				Name:           queueName,
				PriorityFactor: server.queueManagementConfig.DefaultPriorityFactor,
			}
			e := server.queueRepository.CreateQueue(queue)
			if e != nil {
				return status.Errorf(codes.Aborted, e.Error()), []string{}
			}
			return nil, []string{}
		} else {
			return status.Errorf(codes.NotFound, "Queue %q not found", queueName), []string{}
		}
	} else if e != nil {
		return status.Errorf(codes.Unavailable, "Could not load queue %q: %s", queueName, e.Error()), []string{}
	}

	permissionToCheck := basicPermission
	owned, groups := server.permissions.UserOwns(ctx, queue)
	if !owned {
		permissionToCheck = allQueuesPermission
	}
	if e := checkPermission(server.permissions, ctx, permissionToCheck); e != nil {
		return e, []string{}
	}
	return nil, groups
}

func (server *SubmitServer) createJobs(request *api.JobSubmitRequest, owner string, ownershipGroups []string) ([]*api.Job, error) {
	jobs := make([]*api.Job, 0, len(request.JobRequestItems))

	if request.JobSetId == "" {
		return nil, fmt.Errorf("job set is not specified")
	}

	if request.Queue == "" {
		return nil, fmt.Errorf("queue is not specified")
	}

	for i, item := range request.JobRequestItems {

		if item.PodSpec != nil && len(item.PodSpecs) > 0 {
			return nil, fmt.Errorf("job with index %v has both pod spec and pod spec list specified", i)
		}

		if len(item.GetAllPodSpecs()) == 0 {
			return nil, fmt.Errorf("job with index %v has no pod spec", i)
		}

		e := validation.ValidateJobSubmitRequestItem(item)
		if e != nil {
			return nil, fmt.Errorf("job with index %v: %v", i, e)
		}

		namespace := item.Namespace
		if namespace == "" {
			namespace = "default"
		}

		for j, podSpec := range item.GetAllPodSpecs() {
			server.applyDefaultsToPodSpec(podSpec)
			e := validation.ValidatePodSpec(podSpec)
			if e != nil {
				return nil, fmt.Errorf("error validating pod spec of job with index %v, pod: %v: %v", i, j, e)
			}

			// TODO: remove, RequiredNodeLabels is deprecated and will be removed in future versions
			for k, v := range item.RequiredNodeLabels {
				if podSpec.NodeSelector == nil {
					podSpec.NodeSelector = map[string]string{}
				}
				podSpec.NodeSelector[k] = v
			}
		}

		j := &api.Job{
			Id:       util.NewULID(),
			ClientId: item.ClientId,
			Queue:    request.Queue,
			JobSetId: request.JobSetId,

			Namespace:   namespace,
			Labels:      item.Labels,
			Annotations: item.Annotations,

			RequiredNodeLabels: item.RequiredNodeLabels,
			Ingress:            item.Ingress,

			Priority: item.Priority,

			PodSpec:                  item.PodSpec,
			PodSpecs:                 item.PodSpecs,
			Created:                  time.Now(),
			Owner:                    owner,
			QueueOwnershipUserGroups: ownershipGroups,
		}
		jobs = append(jobs, j)
	}

	return jobs, nil
}

func (server *SubmitServer) applyDefaultsToPodSpec(spec *v1.PodSpec) {
	if spec != nil {
		for i := range spec.Containers {
			c := &spec.Containers[i]
			if c.Resources.Limits == nil {
				c.Resources.Limits = map[v1.ResourceName]resource.Quantity{}
			}
			if c.Resources.Requests == nil {
				c.Resources.Requests = map[v1.ResourceName]resource.Quantity{}
			}
			for k, v := range server.defaultJobLimits {
				_, limitExists := c.Resources.Limits[v1.ResourceName(k)]
				_, requestExists := c.Resources.Limits[v1.ResourceName(k)]
				if !limitExists && !requestExists {
					c.Resources.Requests[v1.ResourceName(k)] = v
					c.Resources.Limits[v1.ResourceName(k)] = v
				}
			}
		}
		tolerationsToAdd := []v1.Toleration{}
		for _, defaultToleration := range server.defaultJobTolerations {
			exists := false
			for _, existingToleration := range spec.Tolerations {
				if defaultToleration.MatchToleration(&existingToleration) {
					exists = true
					break
				}
			}
			if !exists {
				tolerationsToAdd = append(tolerationsToAdd, defaultToleration)
			}
		}
		spec.Tolerations = append(spec.Tolerations, tolerationsToAdd...)
	}
}

func validateQueue(queue *api.Queue) error {
	if queue.PriorityFactor < 1.0 {
		return status.Errorf(codes.InvalidArgument, "Minimum queue priority factor is 1.")
	}
	return nil
}
