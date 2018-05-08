// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateTaskRequest struct {
	X                    *wrappers.StringValue `protobuf:"bytes,1,opt,name=_" json:"_,omitempty"`
	JobId                *wrappers.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	NodeId               *wrappers.StringValue `protobuf:"bytes,3,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Target               *wrappers.StringValue `protobuf:"bytes,4,opt,name=target" json:"target,omitempty"`
	TaskAction           *wrappers.StringValue `protobuf:"bytes,5,opt,name=task_action,json=taskAction" json:"task_action,omitempty"`
	Directive            *wrappers.StringValue `protobuf:"bytes,6,opt,name=directive" json:"directive,omitempty"`
	FailureAllowed       *wrappers.BoolValue   `protobuf:"bytes,7,opt,name=failure_allowed,json=failureAllowed" json:"failure_allowed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_f0d9d92ed3ccea9b, []int{0}
}
func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(dst, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetX() *wrappers.StringValue {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *CreateTaskRequest) GetJobId() *wrappers.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *CreateTaskRequest) GetNodeId() *wrappers.StringValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *CreateTaskRequest) GetTarget() *wrappers.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *CreateTaskRequest) GetTaskAction() *wrappers.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *CreateTaskRequest) GetDirective() *wrappers.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *CreateTaskRequest) GetFailureAllowed() *wrappers.BoolValue {
	if m != nil {
		return m.FailureAllowed
	}
	return nil
}

type CreateTaskResponse struct {
	TaskId               *wrappers.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId                *wrappers.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_f0d9d92ed3ccea9b, []int{1}
}
func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(dst, src)
}
func (m *CreateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse.Size(m)
}
func (m *CreateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse proto.InternalMessageInfo

func (m *CreateTaskResponse) GetTaskId() *wrappers.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *CreateTaskResponse) GetJobId() *wrappers.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

type Task struct {
	TaskId               *wrappers.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId                *wrappers.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	TaskAction           *wrappers.StringValue `protobuf:"bytes,3,opt,name=task_action,json=taskAction" json:"task_action,omitempty"`
	Status               *wrappers.StringValue `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	ErrorCode            *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	Directive            *wrappers.StringValue `protobuf:"bytes,6,opt,name=directive" json:"directive,omitempty"`
	Executor             *wrappers.StringValue `protobuf:"bytes,7,opt,name=executor" json:"executor,omitempty"`
	Owner                *wrappers.StringValue `protobuf:"bytes,8,opt,name=owner" json:"owner,omitempty"`
	Target               *wrappers.StringValue `protobuf:"bytes,9,opt,name=target" json:"target,omitempty"`
	NodeId               *wrappers.StringValue `protobuf:"bytes,10,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	CreateTime           *timestamp.Timestamp  `protobuf:"bytes,11,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime           *timestamp.Timestamp  `protobuf:"bytes,12,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
	FailureAllowed       *wrappers.BoolValue   `protobuf:"bytes,13,opt,name=failure_allowed,json=failureAllowed" json:"failure_allowed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_f0d9d92ed3ccea9b, []int{2}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetTaskId() *wrappers.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *Task) GetJobId() *wrappers.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *Task) GetTaskAction() *wrappers.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *Task) GetStatus() *wrappers.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Task) GetErrorCode() *wrappers.UInt32Value {
	if m != nil {
		return m.ErrorCode
	}
	return nil
}

func (m *Task) GetDirective() *wrappers.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *Task) GetExecutor() *wrappers.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *Task) GetOwner() *wrappers.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Task) GetTarget() *wrappers.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Task) GetNodeId() *wrappers.StringValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Task) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Task) GetStatusTime() *timestamp.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

func (m *Task) GetFailureAllowed() *wrappers.BoolValue {
	if m != nil {
		return m.FailureAllowed
	}
	return nil
}

type DescribeTasksRequest struct {
	TaskId   []string              `protobuf:"bytes,1,rep,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId    []string              `protobuf:"bytes,2,rep,name=job_id,json=jobId" json:"job_id,omitempty"`
	Executor *wrappers.StringValue `protobuf:"bytes,3,opt,name=executor" json:"executor,omitempty"`
	Target   *wrappers.StringValue `protobuf:"bytes,4,opt,name=target" json:"target,omitempty"`
	Status   []string              `protobuf:"bytes,5,rep,name=status" json:"status,omitempty"`
	// default is 20, max value is 200
	Limit uint32 `protobuf:"varint,6,opt,name=limit" json:"limit,omitempty"`
	// default is 0
	Offset               uint32   `protobuf:"varint,7,opt,name=offset" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeTasksRequest) Reset()         { *m = DescribeTasksRequest{} }
func (m *DescribeTasksRequest) String() string { return proto.CompactTextString(m) }
func (*DescribeTasksRequest) ProtoMessage()    {}
func (*DescribeTasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_f0d9d92ed3ccea9b, []int{3}
}
func (m *DescribeTasksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeTasksRequest.Unmarshal(m, b)
}
func (m *DescribeTasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeTasksRequest.Marshal(b, m, deterministic)
}
func (dst *DescribeTasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeTasksRequest.Merge(dst, src)
}
func (m *DescribeTasksRequest) XXX_Size() int {
	return xxx_messageInfo_DescribeTasksRequest.Size(m)
}
func (m *DescribeTasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeTasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeTasksRequest proto.InternalMessageInfo

func (m *DescribeTasksRequest) GetTaskId() []string {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *DescribeTasksRequest) GetJobId() []string {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *DescribeTasksRequest) GetExecutor() *wrappers.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *DescribeTasksRequest) GetTarget() *wrappers.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *DescribeTasksRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeTasksRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeTasksRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type DescribeTasksResponse struct {
	TotalCount           uint32   `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	TaskSet              []*Task  `protobuf:"bytes,2,rep,name=task_set,json=taskSet" json:"task_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeTasksResponse) Reset()         { *m = DescribeTasksResponse{} }
func (m *DescribeTasksResponse) String() string { return proto.CompactTextString(m) }
func (*DescribeTasksResponse) ProtoMessage()    {}
func (*DescribeTasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_f0d9d92ed3ccea9b, []int{4}
}
func (m *DescribeTasksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeTasksResponse.Unmarshal(m, b)
}
func (m *DescribeTasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeTasksResponse.Marshal(b, m, deterministic)
}
func (dst *DescribeTasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeTasksResponse.Merge(dst, src)
}
func (m *DescribeTasksResponse) XXX_Size() int {
	return xxx_messageInfo_DescribeTasksResponse.Size(m)
}
func (m *DescribeTasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeTasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeTasksResponse proto.InternalMessageInfo

func (m *DescribeTasksResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeTasksResponse) GetTaskSet() []*Task {
	if m != nil {
		return m.TaskSet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTaskRequest)(nil), "openpitrix.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "openpitrix.CreateTaskResponse")
	proto.RegisterType((*Task)(nil), "openpitrix.Task")
	proto.RegisterType((*DescribeTasksRequest)(nil), "openpitrix.DescribeTasksRequest")
	proto.RegisterType((*DescribeTasksResponse)(nil), "openpitrix.DescribeTasksResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskManager service

type TaskManagerClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	DescribeTasks(ctx context.Context, in *DescribeTasksRequest, opts ...grpc.CallOption) (*DescribeTasksResponse, error)
}

type taskManagerClient struct {
	cc *grpc.ClientConn
}

func NewTaskManagerClient(cc *grpc.ClientConn) TaskManagerClient {
	return &taskManagerClient{cc}
}

func (c *taskManagerClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := grpc.Invoke(ctx, "/openpitrix.TaskManager/CreateTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) DescribeTasks(ctx context.Context, in *DescribeTasksRequest, opts ...grpc.CallOption) (*DescribeTasksResponse, error) {
	out := new(DescribeTasksResponse)
	err := grpc.Invoke(ctx, "/openpitrix.TaskManager/DescribeTasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskManager service

type TaskManagerServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	DescribeTasks(context.Context, *DescribeTasksRequest) (*DescribeTasksResponse, error)
}

func RegisterTaskManagerServer(s *grpc.Server, srv TaskManagerServer) {
	s.RegisterService(&_TaskManager_serviceDesc, srv)
}

func _TaskManager_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_DescribeTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).DescribeTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/DescribeTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).DescribeTasks(ctx, req.(*DescribeTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.TaskManager",
	HandlerType: (*TaskManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskManager_CreateTask_Handler,
		},
		{
			MethodName: "DescribeTasks",
			Handler:    _TaskManager_DescribeTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_task_f0d9d92ed3ccea9b) }

var fileDescriptor_task_f0d9d92ed3ccea9b = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0xdc, 0x34, 0x6e, 0x73, 0xfd, 0x85, 0x9f, 0x51, 0x0b, 0x56, 0x54, 0xda, 0x90, 0x55,
	0x11, 0x34, 0x81, 0xb4, 0x48, 0x88, 0x8a, 0x45, 0x1b, 0x36, 0x11, 0x62, 0x93, 0x16, 0x16, 0x6c,
	0xac, 0x89, 0x7d, 0x63, 0xa6, 0x75, 0x3d, 0x66, 0x66, 0xdc, 0x74, 0x89, 0x58, 0xf0, 0x00, 0xf0,
	0x02, 0xbc, 0x13, 0x8f, 0x00, 0x4f, 0x81, 0x58, 0xa0, 0x19, 0x3b, 0xad, 0xe9, 0x1f, 0x4e, 0x91,
	0x58, 0x45, 0x99, 0x39, 0x67, 0x66, 0xee, 0x3d, 0xe7, 0x5c, 0x03, 0x28, 0x2a, 0xf7, 0xdb, 0x89,
	0xe0, 0x8a, 0x13, 0xe0, 0x09, 0xc6, 0x09, 0x53, 0x82, 0x1d, 0x35, 0x96, 0x43, 0xce, 0xc3, 0x08,
	0x3b, 0x66, 0x67, 0x98, 0x8e, 0x3a, 0x63, 0x41, 0x93, 0x04, 0x85, 0xcc, 0xb0, 0x8d, 0x95, 0xd3,
	0xfb, 0x8a, 0x1d, 0xa0, 0x54, 0xf4, 0x20, 0xc9, 0x01, 0x4b, 0x39, 0x80, 0x26, 0xac, 0x43, 0xe3,
	0x98, 0x2b, 0xaa, 0x18, 0x8f, 0x27, 0xf4, 0x07, 0xe6, 0xc7, 0x5f, 0x0b, 0x31, 0x5e, 0x93, 0x63,
	0x1a, 0x86, 0x28, 0x3a, 0x3c, 0x31, 0x88, 0xb3, 0xe8, 0xd6, 0x97, 0x0a, 0xdc, 0xec, 0x09, 0xa4,
	0x0a, 0x77, 0xa9, 0xdc, 0x1f, 0xe0, 0xbb, 0x14, 0xa5, 0x22, 0xf7, 0xc0, 0xf2, 0x5c, 0xab, 0x69,
	0xad, 0x3a, 0xdd, 0xa5, 0x76, 0x76, 0x5b, 0x7b, 0xf2, 0x9c, 0xf6, 0x8e, 0x12, 0x2c, 0x0e, 0x5f,
	0xd3, 0x28, 0xc5, 0xc1, 0x7f, 0x64, 0x1d, 0xec, 0x3d, 0x3e, 0xf4, 0x58, 0xe0, 0xce, 0x94, 0xc0,
	0x57, 0xf7, 0xf8, 0xb0, 0x1f, 0x90, 0xc7, 0x30, 0x17, 0xf3, 0x00, 0x35, 0xab, 0x52, 0x82, 0x65,
	0x6b, 0x70, 0x3f, 0x20, 0x1b, 0x60, 0x2b, 0x2a, 0x42, 0x54, 0xee, 0x6c, 0x19, 0x56, 0x86, 0x25,
	0xcf, 0xc0, 0xd1, 0x4a, 0x78, 0xd4, 0xd7, 0x85, 0xbb, 0xd5, 0x12, 0x54, 0x23, 0xdd, 0x96, 0xc1,
	0x93, 0xa7, 0x50, 0x0b, 0x98, 0x40, 0x5f, 0xb1, 0x43, 0x74, 0xed, 0x12, 0xe4, 0x13, 0x38, 0xe9,
	0xc1, 0xf5, 0x11, 0x65, 0x51, 0x2a, 0xd0, 0xa3, 0x51, 0xc4, 0xc7, 0x18, 0xb8, 0x73, 0xe6, 0x84,
	0xc6, 0x99, 0x13, 0xb6, 0x39, 0x8f, 0x32, 0xfe, 0xb5, 0x9c, 0xb2, 0x95, 0x31, 0x5a, 0xef, 0x2d,
	0x20, 0x45, 0x89, 0x64, 0xc2, 0x63, 0x89, 0xba, 0x87, 0xa6, 0x2c, 0x16, 0x94, 0x52, 0xca, 0xd6,
	0xe0, 0x7e, 0x70, 0x25, 0xbd, 0x5a, 0x3f, 0xaa, 0x30, 0xab, 0x2f, 0xff, 0x97, 0x97, 0x9e, 0xd6,
	0xad, 0x32, 0xa5, 0x6e, 0x1b, 0x60, 0x4b, 0x45, 0x55, 0x2a, 0xcb, 0x99, 0x25, 0xc3, 0x92, 0x4d,
	0x00, 0x14, 0x82, 0x0b, 0xcf, 0xe7, 0x01, 0x5e, 0xe8, 0x95, 0x57, 0xfd, 0x58, 0xad, 0x77, 0x73,
	0xb9, 0x0d, 0xbe, 0xc7, 0x03, 0xfc, 0x2b, 0xab, 0x3c, 0x81, 0x79, 0x3c, 0x42, 0x3f, 0x55, 0x5c,
	0xe4, 0x1e, 0xb9, 0x9c, 0x7a, 0x8c, 0x26, 0x5d, 0xa8, 0xf2, 0x71, 0x8c, 0xc2, 0x9d, 0x2f, 0xd3,
	0x5b, 0x03, 0x2d, 0x24, 0xa9, 0x36, 0x45, 0x92, 0x0a, 0xb1, 0x85, 0x29, 0x62, 0xbb, 0x09, 0x8e,
	0x6f, 0xfc, 0xeb, 0xe9, 0x49, 0xe6, 0x3a, 0x17, 0x24, 0x60, 0x77, 0x32, 0xe6, 0x06, 0x90, 0xc1,
	0xf5, 0x82, 0x26, 0x67, 0xd2, 0x64, 0xe4, 0xff, 0xff, 0x4c, 0xce, 0xe0, 0x86, 0x7c, 0x4e, 0xfe,
	0xea, 0x53, 0xe7, 0xef, 0xa7, 0x05, 0x0b, 0xcf, 0x51, 0xfa, 0x82, 0x0d, 0x4d, 0x02, 0xe5, 0x64,
	0x4a, 0xde, 0x2e, 0x86, 0xa1, 0xb2, 0x5a, 0x3b, 0xb6, 0xfb, 0x62, 0xc1, 0xee, 0x7a, 0x3d, 0x37,
	0x74, 0x51, 0xe2, 0xca, 0x54, 0x12, 0x5f, 0x6d, 0xf0, 0xdd, 0x3a, 0x4e, 0x40, 0x35, 0x7b, 0x5e,
	0xee, 0xf1, 0x05, 0xa8, 0x46, 0xec, 0x80, 0x29, 0x63, 0xd1, 0xfa, 0x20, 0xfb, 0xa3, 0xd1, 0x7c,
	0x34, 0x92, 0xa8, 0x8c, 0xfd, 0xea, 0x83, 0xfc, 0x5f, 0x0b, 0x61, 0xf1, 0x54, 0xf5, 0xf9, 0x00,
	0x5a, 0x01, 0x47, 0x71, 0x45, 0x23, 0xcf, 0xe7, 0x69, 0xac, 0xcc, 0x3c, 0xa8, 0x0f, 0xc0, 0x2c,
	0xf5, 0xf4, 0x0a, 0xb9, 0x0f, 0xf3, 0xa6, 0x3f, 0xfa, 0x4c, 0xdd, 0x08, 0xa7, 0x7b, 0xa3, 0x7d,
	0xf2, 0x1d, 0x6c, 0x9b, 0x69, 0x66, 0x3a, 0xb8, 0x83, 0xaa, 0xfb, 0xcd, 0x02, 0x47, 0xaf, 0xbc,
	0xa4, 0x31, 0x0d, 0x51, 0x90, 0x17, 0x00, 0x27, 0x43, 0x8f, 0xdc, 0x29, 0x12, 0xcf, 0x7c, 0xaf,
	0x1a, 0xcb, 0x17, 0x6d, 0xe7, 0x4f, 0xfd, 0x68, 0x41, 0xfd, 0xb7, 0x22, 0x48, 0xb3, 0xc8, 0x38,
	0x4f, 0xdd, 0xc6, 0xdd, 0x4b, 0x10, 0xd9, 0xb1, 0xad, 0x87, 0x9f, 0xb6, 0x96, 0x48, 0x23, 0xc8,
	0xf7, 0x9a, 0xba, 0x14, 0xd9, 0x1c, 0x33, 0xf5, 0xb6, 0x39, 0x62, 0x91, 0x42, 0xf1, 0xe1, 0xeb,
	0xf7, 0xcf, 0x33, 0x0e, 0xa9, 0x75, 0x0e, 0x1f, 0x75, 0xcc, 0xe6, 0xf6, 0xec, 0x9b, 0x99, 0x64,
	0x38, 0xb4, 0x8d, 0x6c, 0xeb, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x13, 0xb5, 0x53, 0x22,
	0x08, 0x00, 0x00,
}
