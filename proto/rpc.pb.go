// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PutRequest struct {
	Topic                []byte   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Message              []byte   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Tag                  []byte   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	QueueId              int32    `protobuf:"varint,4,opt,name=queueId,proto3" json:"queueId,omitempty"`
	Delay                int32    `protobuf:"varint,5,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (m *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(m, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetTopic() []byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *PutRequest) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *PutRequest) GetTag() []byte {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *PutRequest) GetQueueId() int32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *PutRequest) GetDelay() int32 {
	if m != nil {
		return m.Delay
	}
	return 0
}

type PutReply struct {
	MessageId            int64    `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutReply) Reset()         { *m = PutReply{} }
func (m *PutReply) String() string { return proto.CompactTextString(m) }
func (*PutReply) ProtoMessage()    {}
func (*PutReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *PutReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutReply.Unmarshal(m, b)
}
func (m *PutReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutReply.Marshal(b, m, deterministic)
}
func (m *PutReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutReply.Merge(m, src)
}
func (m *PutReply) XXX_Size() int {
	return xxx_messageInfo_PutReply.Size(m)
}
func (m *PutReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PutReply.DiscardUnknown(m)
}

var xxx_messageInfo_PutReply proto.InternalMessageInfo

func (m *PutReply) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type PullRequest struct {
	Topic                []byte   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Group                []byte   `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Tag                  []byte   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	QueueId              int32    `protobuf:"varint,4,opt,name=queueId,proto3" json:"queueId,omitempty"`
	PullNum              int32    `protobuf:"varint,5,opt,name=pullNum,proto3" json:"pullNum,omitempty"`
	ReadOffset           int64    `protobuf:"varint,6,opt,name=readOffset,proto3" json:"readOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullRequest) Reset()         { *m = PullRequest{} }
func (m *PullRequest) String() string { return proto.CompactTextString(m) }
func (*PullRequest) ProtoMessage()    {}
func (*PullRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *PullRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullRequest.Unmarshal(m, b)
}
func (m *PullRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullRequest.Marshal(b, m, deterministic)
}
func (m *PullRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullRequest.Merge(m, src)
}
func (m *PullRequest) XXX_Size() int {
	return xxx_messageInfo_PullRequest.Size(m)
}
func (m *PullRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PullRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PullRequest proto.InternalMessageInfo

func (m *PullRequest) GetTopic() []byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *PullRequest) GetGroup() []byte {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *PullRequest) GetTag() []byte {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *PullRequest) GetQueueId() int32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *PullRequest) GetPullNum() int32 {
	if m != nil {
		return m.PullNum
	}
	return 0
}

func (m *PullRequest) GetReadOffset() int64 {
	if m != nil {
		return m.ReadOffset
	}
	return 0
}

type PullReply struct {
	MessageId            int64    `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Message              []byte   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullReply) Reset()         { *m = PullReply{} }
func (m *PullReply) String() string { return proto.CompactTextString(m) }
func (*PullReply) ProtoMessage()    {}
func (*PullReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *PullReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullReply.Unmarshal(m, b)
}
func (m *PullReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullReply.Marshal(b, m, deterministic)
}
func (m *PullReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullReply.Merge(m, src)
}
func (m *PullReply) XXX_Size() int {
	return xxx_messageInfo_PullReply.Size(m)
}
func (m *PullReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PullReply.DiscardUnknown(m)
}

var xxx_messageInfo_PullReply proto.InternalMessageInfo

func (m *PullReply) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *PullReply) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*PutRequest)(nil), "proto.PutRequest")
	proto.RegisterType((*PutReply)(nil), "proto.PutReply")
	proto.RegisterType((*PullRequest)(nil), "proto.PullRequest")
	proto.RegisterType((*PullReply)(nil), "proto.PullReply")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0xb3, 0xb4, 0x05, 0x3b, 0x92, 0x80, 0x13, 0x0e, 0x1b, 0x42, 0x0c, 0xe9, 0xa9, 0xf1,
	0x40, 0x8d, 0xde, 0x38, 0x78, 0xf1, 0xc4, 0x45, 0x49, 0xdf, 0x60, 0xa5, 0x4b, 0x43, 0xb2, 0xb0,
	0x43, 0xbb, 0x7b, 0xe0, 0x62, 0xa2, 0xaf, 0xe0, 0x03, 0xf8, 0x50, 0xbe, 0x82, 0x0f, 0x62, 0xba,
	0xa5, 0x80, 0x07, 0x35, 0x9e, 0x76, 0xff, 0x3f, 0xff, 0xcc, 0x7e, 0x3b, 0x03, 0x61, 0x41, 0x8b,
	0x09, 0x15, 0xda, 0x68, 0x0c, 0xdc, 0x31, 0x1c, 0xe5, 0x5a, 0xe7, 0x4a, 0x26, 0x82, 0x56, 0x89,
	0xd8, 0x6c, 0xb4, 0x11, 0x66, 0xa5, 0x37, 0x65, 0x1d, 0x8a, 0x9e, 0x01, 0xe6, 0xd6, 0xa4, 0x72,
	0x6b, 0x65, 0x69, 0x70, 0x00, 0x81, 0xd1, 0xb4, 0x5a, 0x70, 0x36, 0x66, 0x71, 0x37, 0xad, 0x05,
	0x72, 0xe8, 0xac, 0x65, 0x59, 0x8a, 0x5c, 0xf2, 0x96, 0xf3, 0x1b, 0x89, 0x7d, 0xf0, 0x8c, 0xc8,
	0xb9, 0xe7, 0xdc, 0xea, 0x5a, 0x65, 0xb7, 0x56, 0x5a, 0x39, 0xcb, 0xb8, 0x3f, 0x66, 0x71, 0x90,
	0x36, 0xb2, 0xea, 0x9d, 0x49, 0x25, 0x76, 0x3c, 0x70, 0x7e, 0x2d, 0xa2, 0x18, 0xce, 0xdc, 0xfb,
	0xa4, 0x76, 0x38, 0x82, 0x70, 0xdf, 0x78, 0x96, 0x39, 0x02, 0x2f, 0x3d, 0x1a, 0xd1, 0x3b, 0x83,
	0xf3, 0xb9, 0x55, 0xea, 0x77, 0xd6, 0x01, 0x04, 0x79, 0xa1, 0x2d, 0xed, 0x49, 0x6b, 0xf1, 0x2f,
	0x4e, 0x0e, 0x1d, 0xb2, 0x4a, 0x3d, 0xd8, 0xf5, 0x9e, 0xb4, 0x91, 0x78, 0x09, 0x50, 0x48, 0x91,
	0x3d, 0x2e, 0x97, 0xa5, 0x34, 0xbc, 0xed, 0x00, 0x4f, 0x9c, 0xe8, 0x1e, 0xc2, 0x1a, 0xf0, 0xcf,
	0xcf, 0xfc, 0x3c, 0xd2, 0x9b, 0x17, 0x06, 0x5e, 0x4a, 0x0b, 0x9c, 0x82, 0x37, 0xb7, 0x06, 0x2f,
	0xea, 0x3d, 0x4d, 0x8e, 0x4b, 0x1a, 0xf6, 0x4e, 0x2d, 0x52, 0xbb, 0xa8, 0xf7, 0xfa, 0xf1, 0xf9,
	0xd6, 0x0a, 0x23, 0x3f, 0x21, 0x6b, 0xa6, 0xec, 0x0a, 0xef, 0xc0, 0xaf, 0x40, 0x10, 0x0f, 0xc9,
	0xc3, 0xd8, 0x86, 0xfd, 0x6f, 0x5e, 0x55, 0xde, 0x75, 0xe5, 0x6d, 0xf4, 0x13, 0xd2, 0x14, 0xb3,
	0x6b, 0xf6, 0xd4, 0x76, 0x91, 0xdb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x7d, 0xad, 0x21,
	0x4d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RpcClient interface {
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutReply, error)
	Pull(ctx context.Context, opts ...grpc.CallOption) (Rpc_PullClient, error)
}

type rpcClient struct {
	cc *grpc.ClientConn
}

func NewRpcClient(cc *grpc.ClientConn) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutReply, error) {
	out := new(PutReply)
	err := c.cc.Invoke(ctx, "/proto.Rpc/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Pull(ctx context.Context, opts ...grpc.CallOption) (Rpc_PullClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Rpc_serviceDesc.Streams[0], "/proto.Rpc/Pull", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcPullClient{stream}
	return x, nil
}

type Rpc_PullClient interface {
	Send(*PullRequest) error
	Recv() (*PullReply, error)
	grpc.ClientStream
}

type rpcPullClient struct {
	grpc.ClientStream
}

func (x *rpcPullClient) Send(m *PullRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rpcPullClient) Recv() (*PullReply, error) {
	m := new(PullReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RpcServer is the server API for Rpc service.
type RpcServer interface {
	Put(context.Context, *PutRequest) (*PutReply, error)
	Pull(Rpc_PullServer) error
}

// UnimplementedRpcServer can be embedded to have forward compatible implementations.
type UnimplementedRpcServer struct {
}

func (*UnimplementedRpcServer) Put(ctx context.Context, req *PutRequest) (*PutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedRpcServer) Pull(srv Rpc_PullServer) error {
	return status.Errorf(codes.Unimplemented, "method Pull not implemented")
}

func RegisterRpcServer(s *grpc.Server, srv RpcServer) {
	s.RegisterService(&_Rpc_serviceDesc, srv)
}

func _Rpc_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Pull_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RpcServer).Pull(&rpcPullServer{stream})
}

type Rpc_PullServer interface {
	Send(*PullReply) error
	Recv() (*PullRequest, error)
	grpc.ServerStream
}

type rpcPullServer struct {
	grpc.ServerStream
}

func (x *rpcPullServer) Send(m *PullReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rpcPullServer) Recv() (*PullRequest, error) {
	m := new(PullRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Rpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _Rpc_Put_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pull",
			Handler:       _Rpc_Pull_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpc.proto",
}