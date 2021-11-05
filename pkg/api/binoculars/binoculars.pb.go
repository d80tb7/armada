// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/api/binoculars/binoculars.proto

package binoculars

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// swagger:model
type LogRequest struct {
	JobId        string            `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"jobId,omitempty"`
	PodNumber    int32             `protobuf:"varint,2,opt,name=pod_number,json=podNumber,proto3" json:"podNumber,omitempty"`
	PodNamespace string            `protobuf:"bytes,3,opt,name=pod_namespace,json=podNamespace,proto3" json:"podNamespace,omitempty"`
	SinceTime    string            `protobuf:"bytes,4,opt,name=since_time,json=sinceTime,proto3" json:"sinceTime,omitempty"`
	LogOptions   *v1.PodLogOptions `protobuf:"bytes,5,opt,name=log_options,json=logOptions,proto3" json:"logOptions,omitempty"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f2fc8093f6f091f, []int{0}
}
func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(m, src)
}
func (m *LogRequest) XXX_Size() int {
	return m.Size()
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *LogRequest) GetPodNumber() int32 {
	if m != nil {
		return m.PodNumber
	}
	return 0
}

func (m *LogRequest) GetPodNamespace() string {
	if m != nil {
		return m.PodNamespace
	}
	return ""
}

func (m *LogRequest) GetSinceTime() string {
	if m != nil {
		return m.SinceTime
	}
	return ""
}

func (m *LogRequest) GetLogOptions() *v1.PodLogOptions {
	if m != nil {
		return m.LogOptions
	}
	return nil
}

// swagger:model
type LogResponse struct {
	Log string `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (m *LogResponse) Reset()         { *m = LogResponse{} }
func (m *LogResponse) String() string { return proto.CompactTextString(m) }
func (*LogResponse) ProtoMessage()    {}
func (*LogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f2fc8093f6f091f, []int{1}
}
func (m *LogResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LogResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogResponse.Merge(m, src)
}
func (m *LogResponse) XXX_Size() int {
	return m.Size()
}
func (m *LogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogResponse proto.InternalMessageInfo

func (m *LogResponse) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func init() {
	proto.RegisterType((*LogRequest)(nil), "binoculars.LogRequest")
	proto.RegisterType((*LogResponse)(nil), "binoculars.LogResponse")
}

func init() {
	proto.RegisterFile("pkg/api/binoculars/binoculars.proto", fileDescriptor_3f2fc8093f6f091f)
}

var fileDescriptor_3f2fc8093f6f091f = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x6a, 0xe3, 0x30,
	0x18, 0xc4, 0xa3, 0xfc, 0x83, 0x28, 0xbb, 0xb0, 0x08, 0x76, 0x63, 0xc2, 0xc6, 0x9b, 0x75, 0x2e,
	0x61, 0x0f, 0x32, 0xc9, 0x5e, 0x4a, 0x8f, 0xb9, 0x15, 0x42, 0x5b, 0x4c, 0xe9, 0xd5, 0xc8, 0xb6,
	0x10, 0x4a, 0x6c, 0x7d, 0xaa, 0xe5, 0xe4, 0x01, 0xfa, 0x04, 0x85, 0xbe, 0x54, 0xe9, 0x29, 0xd0,
	0x4b, 0x8f, 0x25, 0xe9, 0x83, 0x14, 0xcb, 0x69, 0x12, 0xe8, 0x6d, 0x18, 0xcf, 0x37, 0x78, 0x7e,
	0xc2, 0x23, 0xbd, 0x14, 0x3e, 0xd3, 0xd2, 0x8f, 0xa4, 0x82, 0x78, 0x95, 0xb2, 0xdc, 0x9c, 0x48,
	0xaa, 0x73, 0x28, 0x80, 0xe0, 0xa3, 0xd3, 0xf7, 0x96, 0x67, 0x86, 0x4a, 0xb0, 0x37, 0x31, 0xe4,
	0xdc, 0x5f, 0x4f, 0x7c, 0xc1, 0x15, 0xcf, 0x59, 0xc1, 0x93, 0x2a, 0xdf, 0xff, 0x2d, 0x00, 0x44,
	0xca, 0x6d, 0x86, 0x29, 0x05, 0x05, 0x2b, 0x24, 0xa8, 0x7d, 0x9b, 0xf7, 0x8c, 0x30, 0x9e, 0x83,
	0x08, 0xf8, 0xdd, 0x8a, 0x9b, 0x82, 0xfc, 0xc4, 0xed, 0x05, 0x44, 0xa1, 0x4c, 0x1c, 0x34, 0x44,
	0xe3, 0x4e, 0xd0, 0x5a, 0x40, 0x74, 0x91, 0x90, 0x01, 0xc6, 0x1a, 0x92, 0x50, 0xad, 0xb2, 0x88,
	0xe7, 0x4e, 0x7d, 0x88, 0xc6, 0xad, 0xa0, 0xa3, 0x21, 0xb9, 0xb4, 0x06, 0x19, 0xe1, 0xef, 0xf6,
	0x33, 0xcb, 0xb8, 0xd1, 0x2c, 0xe6, 0x4e, 0xc3, 0x1e, 0x7f, 0x2b, 0x13, 0x9f, 0x5e, 0xd9, 0x61,
	0xa4, 0x8a, 0x79, 0x58, 0xc8, 0x8c, 0x3b, 0x4d, 0x9b, 0xe8, 0x58, 0xe7, 0x46, 0x66, 0x9c, 0xcc,
	0x70, 0x37, 0x05, 0x11, 0x82, 0xb6, 0x7f, 0xe7, 0xb4, 0x86, 0x68, 0xdc, 0x9d, 0xfe, 0xa5, 0xd5,
	0x40, 0xca, 0xb4, 0xa4, 0xe5, 0x40, 0xba, 0x9e, 0xd0, 0x6b, 0x48, 0xe6, 0x20, 0xae, 0xaa, 0x60,
	0x80, 0xd3, 0x83, 0xf6, 0xfe, 0xe0, 0xae, 0xdd, 0x62, 0x34, 0x28, 0xc3, 0xc9, 0x0f, 0xdc, 0x48,
	0x41, 0xec, 0x97, 0x94, 0x72, 0x9a, 0x60, 0x3c, 0x3b, 0xd0, 0x23, 0xb7, 0xb8, 0x39, 0x07, 0x61,
	0xc8, 0x2f, 0x7a, 0x02, 0xf9, 0x08, 0xa3, 0xdf, 0xfb, 0xe2, 0x57, 0xc5, 0xde, 0xe0, 0xfe, 0xe5,
	0xfd, 0xb1, 0xde, 0xf3, 0x48, 0xc9, 0xfb, 0xe4, 0xad, 0x52, 0x10, 0xe7, 0xe8, 0xdf, 0xcc, 0x79,
	0xda, 0xba, 0x68, 0xb3, 0x75, 0xd1, 0xdb, 0xd6, 0x45, 0x0f, 0x3b, 0xb7, 0xb6, 0xd9, 0xb9, 0xb5,
	0xd7, 0x9d, 0x5b, 0x8b, 0xda, 0x16, 0xfa, 0xff, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x11,
	0xac, 0x4c, 0xe9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BinocularsClient is the client API for Binoculars service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BinocularsClient interface {
	Logs(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
}

type binocularsClient struct {
	cc *grpc.ClientConn
}

func NewBinocularsClient(cc *grpc.ClientConn) BinocularsClient {
	return &binocularsClient{cc}
}

func (c *binocularsClient) Logs(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/binoculars.Binoculars/Logs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BinocularsServer is the server API for Binoculars service.
type BinocularsServer interface {
	Logs(context.Context, *LogRequest) (*LogResponse, error)
}

// UnimplementedBinocularsServer can be embedded to have forward compatible implementations.
type UnimplementedBinocularsServer struct {
}

func (*UnimplementedBinocularsServer) Logs(ctx context.Context, req *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logs not implemented")
}

func RegisterBinocularsServer(s *grpc.Server, srv BinocularsServer) {
	s.RegisterService(&_Binoculars_serviceDesc, srv)
}

func _Binoculars_Logs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinocularsServer).Logs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binoculars.Binoculars/Logs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinocularsServer).Logs(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Binoculars_serviceDesc = grpc.ServiceDesc{
	ServiceName: "binoculars.Binoculars",
	HandlerType: (*BinocularsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Logs",
			Handler:    _Binoculars_Logs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/binoculars/binoculars.proto",
}

func (m *LogRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LogRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LogOptions != nil {
		{
			size, err := m.LogOptions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBinoculars(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SinceTime) > 0 {
		i -= len(m.SinceTime)
		copy(dAtA[i:], m.SinceTime)
		i = encodeVarintBinoculars(dAtA, i, uint64(len(m.SinceTime)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PodNamespace) > 0 {
		i -= len(m.PodNamespace)
		copy(dAtA[i:], m.PodNamespace)
		i = encodeVarintBinoculars(dAtA, i, uint64(len(m.PodNamespace)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PodNumber != 0 {
		i = encodeVarintBinoculars(dAtA, i, uint64(m.PodNumber))
		i--
		dAtA[i] = 0x10
	}
	if len(m.JobId) > 0 {
		i -= len(m.JobId)
		copy(dAtA[i:], m.JobId)
		i = encodeVarintBinoculars(dAtA, i, uint64(len(m.JobId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LogResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LogResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintBinoculars(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBinoculars(dAtA []byte, offset int, v uint64) int {
	offset -= sovBinoculars(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LogRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JobId)
	if l > 0 {
		n += 1 + l + sovBinoculars(uint64(l))
	}
	if m.PodNumber != 0 {
		n += 1 + sovBinoculars(uint64(m.PodNumber))
	}
	l = len(m.PodNamespace)
	if l > 0 {
		n += 1 + l + sovBinoculars(uint64(l))
	}
	l = len(m.SinceTime)
	if l > 0 {
		n += 1 + l + sovBinoculars(uint64(l))
	}
	if m.LogOptions != nil {
		l = m.LogOptions.Size()
		n += 1 + l + sovBinoculars(uint64(l))
	}
	return n
}

func (m *LogResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovBinoculars(uint64(l))
	}
	return n
}

func sovBinoculars(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBinoculars(x uint64) (n int) {
	return sovBinoculars(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LogRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBinoculars
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinoculars
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBinoculars
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBinoculars
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodNumber", wireType)
			}
			m.PodNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinoculars
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PodNumber |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodNamespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinoculars
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBinoculars
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBinoculars
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PodNamespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SinceTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinoculars
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBinoculars
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBinoculars
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SinceTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinoculars
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBinoculars
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBinoculars
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LogOptions == nil {
				m.LogOptions = &v1.PodLogOptions{}
			}
			if err := m.LogOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBinoculars(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBinoculars
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LogResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBinoculars
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinoculars
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBinoculars
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBinoculars
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBinoculars(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBinoculars
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBinoculars(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBinoculars
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBinoculars
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBinoculars
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthBinoculars
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBinoculars
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBinoculars
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBinoculars        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBinoculars          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBinoculars = fmt.Errorf("proto: unexpected end of group")
)
