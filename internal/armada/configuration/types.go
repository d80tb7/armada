package configuration

import (
	"time"

	"github.com/go-redis/redis"
	v1 "k8s.io/api/core/v1"

	"github.com/G-Research/armada/internal/common"
	"github.com/G-Research/armada/internal/common/auth/configuration"
)

type ArmadaConfig struct {
	Auth configuration.AuthConfig

	GrpcPort           uint16
	HttpPort           uint16
	MetricsPort        uint16
	CorsAllowedOrigins []string

	PriorityHalfTime time.Duration
	Redis            redis.UniversalOptions
	EventsNats       NatsConfig
	EventsRedis      redis.UniversalOptions

	Scheduling        SchedulingConfig
	QueueManagement   QueueManagementConfig
	DatabaseRetention DatabaseRetentionPolicy
	EventRetention    EventRetentionPolicy

	Metrics MetricsConfig
}

type SchedulingConfig struct {
	UseProbabilisticSchedulingForAllResources bool
	QueueLeaseBatchSize                       uint
	MinimumResourceToSchedule                 common.ComputeResourcesFloat
	MaximalClusterFractionToSchedule          map[string]float64
	MaximalResourceFractionToSchedulePerQueue map[string]float64
	MaximalResourceFractionPerQueue           map[string]float64
	Lease                                     LeaseSettings
	DefaultJobLimits                          common.ComputeResources
	DefaultJobTolerations                     []v1.Toleration
	MaxRetries                                uint // Maximum number of retries before a Job is failed
	ResourceScarcity                          map[string]float64
	PoolResourceScarcity                      map[string]map[string]float64
}

type DatabaseRetentionPolicy struct {
	JobRetentionDuration time.Duration
}

type EventRetentionPolicy struct {
	ExpiryEnabled     bool
	RetentionDuration time.Duration
}

type LeaseSettings struct {
	ExpireAfter        time.Duration
	ExpiryLoopInterval time.Duration
}

type NatsConfig struct {
	Servers        []string
	ClusterID      string
	Subject        string
	QueueGroup     string
	JobStatusGroup string
}

type QueueManagementConfig struct {
	AutoCreateQueues      bool
	DefaultPriorityFactor float64
}

type MetricsConfig struct {
	RefreshInterval time.Duration
}
