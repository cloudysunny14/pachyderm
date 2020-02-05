// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/worker/server/service.proto

package server

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	pps "github.com/pachyderm/pachyderm/src/client/pps"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type CancelRequest struct {
	JobID                string   `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	DataFilters          []string `protobuf:"bytes,1,rep,name=data_filters,json=dataFilters,proto3" json:"data_filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelRequest) Reset()         { *m = CancelRequest{} }
func (m *CancelRequest) String() string { return proto.CompactTextString(m) }
func (*CancelRequest) ProtoMessage()    {}
func (*CancelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4407c0c45dc0204, []int{0}
}
func (m *CancelRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CancelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CancelRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CancelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelRequest.Merge(m, src)
}
func (m *CancelRequest) XXX_Size() int {
	return m.Size()
}
func (m *CancelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelRequest proto.InternalMessageInfo

func (m *CancelRequest) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *CancelRequest) GetDataFilters() []string {
	if m != nil {
		return m.DataFilters
	}
	return nil
}

type CancelResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelResponse) Reset()         { *m = CancelResponse{} }
func (m *CancelResponse) String() string { return proto.CompactTextString(m) }
func (*CancelResponse) ProtoMessage()    {}
func (*CancelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4407c0c45dc0204, []int{1}
}
func (m *CancelResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CancelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CancelResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CancelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelResponse.Merge(m, src)
}
func (m *CancelResponse) XXX_Size() int {
	return m.Size()
}
func (m *CancelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CancelResponse proto.InternalMessageInfo

func (m *CancelResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type GetChunkRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Shard                int64    `protobuf:"varint,2,opt,name=shard,proto3" json:"shard,omitempty"`
	Stats                bool     `protobuf:"varint,3,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetChunkRequest) Reset()         { *m = GetChunkRequest{} }
func (m *GetChunkRequest) String() string { return proto.CompactTextString(m) }
func (*GetChunkRequest) ProtoMessage()    {}
func (*GetChunkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4407c0c45dc0204, []int{2}
}
func (m *GetChunkRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetChunkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetChunkRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetChunkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChunkRequest.Merge(m, src)
}
func (m *GetChunkRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetChunkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChunkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetChunkRequest proto.InternalMessageInfo

func (m *GetChunkRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetChunkRequest) GetShard() int64 {
	if m != nil {
		return m.Shard
	}
	return 0
}

func (m *GetChunkRequest) GetStats() bool {
	if m != nil {
		return m.Stats
	}
	return false
}

func init() {
	proto.RegisterType((*CancelRequest)(nil), "server.CancelRequest")
	proto.RegisterType((*CancelResponse)(nil), "server.CancelResponse")
	proto.RegisterType((*GetChunkRequest)(nil), "server.GetChunkRequest")
}

func init() { proto.RegisterFile("server/worker/server/service.proto", fileDescriptor_c4407c0c45dc0204) }

var fileDescriptor_c4407c0c45dc0204 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0xed, 0x34, 0x34, 0x6e, 0x47, 0x5d, 0x71, 0xa8, 0x6b, 0xc8, 0x42, 0xac, 0x39, 0x15, 0x0f,
	0x33, 0xa2, 0x88, 0x78, 0xed, 0xee, 0x2a, 0x2b, 0x78, 0x89, 0xa2, 0xe0, 0x65, 0x99, 0x4c, 0xbe,
	0x4d, 0xb2, 0x4d, 0x33, 0xe3, 0xcc, 0xc4, 0xd2, 0x7f, 0xe7, 0xd1, 0xa3, 0xbf, 0x40, 0x24, 0xbf,
	0x44, 0x92, 0x49, 0x40, 0xab, 0x87, 0x90, 0xef, 0xbd, 0xef, 0xf1, 0x78, 0xdf, 0x4b, 0x70, 0x6c,
	0x40, 0x7f, 0x05, 0xcd, 0x76, 0x52, 0x6f, 0x40, 0xb3, 0x01, 0x75, 0xaf, 0x52, 0x00, 0x55, 0x5a,
	0x5a, 0x49, 0x7c, 0xc7, 0x86, 0x0b, 0x51, 0x95, 0x50, 0x5b, 0xa6, 0x94, 0xe9, 0x1e, 0xb7, 0x0d,
	0x17, 0xb9, 0xcc, 0x65, 0x3f, 0xb2, 0x6e, 0x1a, 0xd8, 0xd3, 0x5c, 0xca, 0xbc, 0x02, 0xd6, 0xa3,
	0xb4, 0xb9, 0x66, 0xb0, 0x55, 0x76, 0x3f, 0x2c, 0xa3, 0xc3, 0xe5, 0x4e, 0x73, 0xa5, 0x40, 0x0f,
	0x96, 0xf1, 0x07, 0x7c, 0xf7, 0x8c, 0xd7, 0x02, 0xaa, 0x04, 0xbe, 0x34, 0x60, 0x2c, 0x59, 0x62,
	0xff, 0x46, 0xa6, 0x57, 0x65, 0x16, 0x4c, 0x97, 0x68, 0x35, 0x5f, 0xcf, 0xdb, 0x9f, 0x8f, 0x66,
	0x6f, 0x65, 0x7a, 0x79, 0x9e, 0xcc, 0x6e, 0x64, 0x7a, 0x99, 0x91, 0xc7, 0xf8, 0x4e, 0xc6, 0x2d,
	0xbf, 0xba, 0x2e, 0x2b, 0x0b, 0xda, 0x04, 0x68, 0xe9, 0xad, 0xe6, 0xc9, 0xed, 0x8e, 0x7b, 0xed,
	0xa8, 0xf8, 0x09, 0x3e, 0x1e, 0x5d, 0x8d, 0x92, 0xb5, 0x01, 0x12, 0xe0, 0x5b, 0xa6, 0x11, 0x02,
	0x4c, 0xa7, 0x47, 0xab, 0xa3, 0x64, 0x84, 0xf1, 0x3b, 0x7c, 0xef, 0x0d, 0xd8, 0xb3, 0xa2, 0xa9,
	0x37, 0x63, 0x86, 0x63, 0x3c, 0x2d, 0xb3, 0x5e, 0xe7, 0x25, 0xd3, 0x32, 0x23, 0x0b, 0x3c, 0x33,
	0x05, 0xd7, 0x2e, 0x92, 0x97, 0x38, 0xd0, 0xb3, 0x96, 0x5b, 0x13, 0x78, 0xbd, 0xa1, 0x03, 0xcf,
	0xbe, 0x21, 0xec, 0x7f, 0xea, 0x1b, 0x26, 0x2f, 0xb0, 0xff, 0xde, 0x72, 0xdb, 0x18, 0x72, 0x42,
	0x5d, 0x0d, 0x74, 0xac, 0x81, 0x5e, 0x74, 0x1d, 0x85, 0xf7, 0x69, 0x57, 0xae, 0x93, 0x3b, 0x69,
	0x3c, 0x21, 0xaf, 0xb0, 0xef, 0xc2, 0x93, 0x07, 0xd4, 0x7d, 0x0e, 0xfa, 0x57, 0x45, 0xe1, 0xc9,
	0x21, 0xed, 0x6e, 0x8c, 0x27, 0xe4, 0x1c, 0x1f, 0x8d, 0xb7, 0x90, 0x87, 0xa3, 0xea, 0xe0, 0xba,
	0xf0, 0xf4, 0x9f, 0x30, 0xeb, 0xbd, 0x05, 0xf3, 0x91, 0x57, 0x0d, 0xc4, 0x93, 0xa7, 0x68, 0x7d,
	0xf1, 0xbd, 0x8d, 0xd0, 0x8f, 0x36, 0x42, 0xbf, 0xda, 0x08, 0x7d, 0x7e, 0x99, 0x97, 0xb6, 0x68,
	0x52, 0x2a, 0xe4, 0x96, 0x29, 0x2e, 0x8a, 0x7d, 0x06, 0xfa, 0xcf, 0xc9, 0x68, 0xc1, 0xfe, 0xf7,
	0x67, 0xa5, 0x7e, 0xef, 0xff, 0xfc, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x94, 0x34, 0x88,
	0x78, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerClient interface {
	Status(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*pps.WorkerStatus, error)
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
	GetChunk(ctx context.Context, in *GetChunkRequest, opts ...grpc.CallOption) (Worker_GetChunkClient, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) Status(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*pps.WorkerStatus, error) {
	out := new(pps.WorkerStatus)
	err := c.cc.Invoke(ctx, "/server.Worker/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := c.cc.Invoke(ctx, "/server.Worker/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) GetChunk(ctx context.Context, in *GetChunkRequest, opts ...grpc.CallOption) (Worker_GetChunkClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Worker_serviceDesc.Streams[0], "/server.Worker/GetChunk", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerGetChunkClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Worker_GetChunkClient interface {
	Recv() (*types.BytesValue, error)
	grpc.ClientStream
}

type workerGetChunkClient struct {
	grpc.ClientStream
}

func (x *workerGetChunkClient) Recv() (*types.BytesValue, error) {
	m := new(types.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WorkerServer is the server API for Worker service.
type WorkerServer interface {
	Status(context.Context, *types.Empty) (*pps.WorkerStatus, error)
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
	GetChunk(*GetChunkRequest, Worker_GetChunkServer) error
}

// UnimplementedWorkerServer can be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (*UnimplementedWorkerServer) Status(ctx context.Context, req *types.Empty) (*pps.WorkerStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedWorkerServer) Cancel(ctx context.Context, req *CancelRequest) (*CancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (*UnimplementedWorkerServer) GetChunk(req *GetChunkRequest, srv Worker_GetChunkServer) error {
	return status.Errorf(codes.Unimplemented, "method GetChunk not implemented")
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Worker/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Status(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Worker/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_GetChunk_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetChunkRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WorkerServer).GetChunk(m, &workerGetChunkServer{stream})
}

type Worker_GetChunkServer interface {
	Send(*types.BytesValue) error
	grpc.ServerStream
}

type workerGetChunkServer struct {
	grpc.ServerStream
}

func (x *workerGetChunkServer) Send(m *types.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Worker_Status_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Worker_Cancel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetChunk",
			Handler:       _Worker_GetChunk_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server/worker/server/service.proto",
}

func (m *CancelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CancelRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.JobID) > 0 {
		i -= len(m.JobID)
		copy(dAtA[i:], m.JobID)
		i = encodeVarintService(dAtA, i, uint64(len(m.JobID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DataFilters) > 0 {
		for iNdEx := len(m.DataFilters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DataFilters[iNdEx])
			copy(dAtA[i:], m.DataFilters[iNdEx])
			i = encodeVarintService(dAtA, i, uint64(len(m.DataFilters[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CancelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CancelResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetChunkRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetChunkRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetChunkRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Stats {
		i--
		if m.Stats {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Shard != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Shard))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CancelRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DataFilters) > 0 {
		for _, s := range m.DataFilters {
			l = len(s)
			n += 1 + l + sovService(uint64(l))
		}
	}
	l = len(m.JobID)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CancelResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetChunkRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovService(uint64(m.Id))
	}
	if m.Shard != 0 {
		n += 1 + sovService(uint64(m.Shard))
	}
	if m.Stats {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CancelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: CancelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataFilters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataFilters = append(m.DataFilters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CancelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: CancelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetChunkRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: GetChunkRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetChunkRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shard", wireType)
			}
			m.Shard = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Shard |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Stats = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)