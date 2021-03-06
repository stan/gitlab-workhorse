// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cleanup.proto

package gitalypb // import "gitlab.com/gitlab-org/gitaly-proto/go/gitalypb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type ApplyBfgObjectMapRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// A raw object-map file as generated by BFG: https://rtyley.github.io/bfg-repo-cleaner
	// Each line in the file has two object SHAs, space-separated - the original
	// SHA of the object, and the SHA after BFG has rewritten the object.
	ObjectMap            []byte   `protobuf:"bytes,2,opt,name=object_map,json=objectMap,proto3" json:"object_map,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyBfgObjectMapRequest) Reset()         { *m = ApplyBfgObjectMapRequest{} }
func (m *ApplyBfgObjectMapRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyBfgObjectMapRequest) ProtoMessage()    {}
func (*ApplyBfgObjectMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cleanup_048c113e3f69de1a, []int{0}
}
func (m *ApplyBfgObjectMapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyBfgObjectMapRequest.Unmarshal(m, b)
}
func (m *ApplyBfgObjectMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyBfgObjectMapRequest.Marshal(b, m, deterministic)
}
func (dst *ApplyBfgObjectMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyBfgObjectMapRequest.Merge(dst, src)
}
func (m *ApplyBfgObjectMapRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyBfgObjectMapRequest.Size(m)
}
func (m *ApplyBfgObjectMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyBfgObjectMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyBfgObjectMapRequest proto.InternalMessageInfo

func (m *ApplyBfgObjectMapRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *ApplyBfgObjectMapRequest) GetObjectMap() []byte {
	if m != nil {
		return m.ObjectMap
	}
	return nil
}

type ApplyBfgObjectMapResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyBfgObjectMapResponse) Reset()         { *m = ApplyBfgObjectMapResponse{} }
func (m *ApplyBfgObjectMapResponse) String() string { return proto.CompactTextString(m) }
func (*ApplyBfgObjectMapResponse) ProtoMessage()    {}
func (*ApplyBfgObjectMapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cleanup_048c113e3f69de1a, []int{1}
}
func (m *ApplyBfgObjectMapResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyBfgObjectMapResponse.Unmarshal(m, b)
}
func (m *ApplyBfgObjectMapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyBfgObjectMapResponse.Marshal(b, m, deterministic)
}
func (dst *ApplyBfgObjectMapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyBfgObjectMapResponse.Merge(dst, src)
}
func (m *ApplyBfgObjectMapResponse) XXX_Size() int {
	return xxx_messageInfo_ApplyBfgObjectMapResponse.Size(m)
}
func (m *ApplyBfgObjectMapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyBfgObjectMapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyBfgObjectMapResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ApplyBfgObjectMapRequest)(nil), "gitaly.ApplyBfgObjectMapRequest")
	proto.RegisterType((*ApplyBfgObjectMapResponse)(nil), "gitaly.ApplyBfgObjectMapResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CleanupServiceClient is the client API for CleanupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CleanupServiceClient interface {
	ApplyBfgObjectMap(ctx context.Context, opts ...grpc.CallOption) (CleanupService_ApplyBfgObjectMapClient, error)
}

type cleanupServiceClient struct {
	cc *grpc.ClientConn
}

func NewCleanupServiceClient(cc *grpc.ClientConn) CleanupServiceClient {
	return &cleanupServiceClient{cc}
}

func (c *cleanupServiceClient) ApplyBfgObjectMap(ctx context.Context, opts ...grpc.CallOption) (CleanupService_ApplyBfgObjectMapClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CleanupService_serviceDesc.Streams[0], "/gitaly.CleanupService/ApplyBfgObjectMap", opts...)
	if err != nil {
		return nil, err
	}
	x := &cleanupServiceApplyBfgObjectMapClient{stream}
	return x, nil
}

type CleanupService_ApplyBfgObjectMapClient interface {
	Send(*ApplyBfgObjectMapRequest) error
	CloseAndRecv() (*ApplyBfgObjectMapResponse, error)
	grpc.ClientStream
}

type cleanupServiceApplyBfgObjectMapClient struct {
	grpc.ClientStream
}

func (x *cleanupServiceApplyBfgObjectMapClient) Send(m *ApplyBfgObjectMapRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cleanupServiceApplyBfgObjectMapClient) CloseAndRecv() (*ApplyBfgObjectMapResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ApplyBfgObjectMapResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CleanupServiceServer is the server API for CleanupService service.
type CleanupServiceServer interface {
	ApplyBfgObjectMap(CleanupService_ApplyBfgObjectMapServer) error
}

func RegisterCleanupServiceServer(s *grpc.Server, srv CleanupServiceServer) {
	s.RegisterService(&_CleanupService_serviceDesc, srv)
}

func _CleanupService_ApplyBfgObjectMap_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CleanupServiceServer).ApplyBfgObjectMap(&cleanupServiceApplyBfgObjectMapServer{stream})
}

type CleanupService_ApplyBfgObjectMapServer interface {
	SendAndClose(*ApplyBfgObjectMapResponse) error
	Recv() (*ApplyBfgObjectMapRequest, error)
	grpc.ServerStream
}

type cleanupServiceApplyBfgObjectMapServer struct {
	grpc.ServerStream
}

func (x *cleanupServiceApplyBfgObjectMapServer) SendAndClose(m *ApplyBfgObjectMapResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cleanupServiceApplyBfgObjectMapServer) Recv() (*ApplyBfgObjectMapRequest, error) {
	m := new(ApplyBfgObjectMapRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CleanupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.CleanupService",
	HandlerType: (*CleanupServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ApplyBfgObjectMap",
			Handler:       _CleanupService_ApplyBfgObjectMap_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "cleanup.proto",
}

func init() { proto.RegisterFile("cleanup.proto", fileDescriptor_cleanup_048c113e3f69de1a) }

var fileDescriptor_cleanup_048c113e3f69de1a = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x49, 0x4d,
	0xcc, 0x2b, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xcf, 0x2c, 0x49, 0xcc,
	0xa9, 0x94, 0xe2, 0x29, 0xce, 0x48, 0x2c, 0x4a, 0x4d, 0x81, 0x88, 0x2a, 0x95, 0x72, 0x49, 0x38,
	0x16, 0x14, 0xe4, 0x54, 0x3a, 0xa5, 0xa5, 0xfb, 0x27, 0x65, 0xa5, 0x26, 0x97, 0xf8, 0x26, 0x16,
	0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x19, 0x71, 0x71, 0x15, 0xa5, 0x16, 0xe4, 0x17,
	0x67, 0x96, 0xe4, 0x17, 0x55, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x09, 0xe9, 0x41, 0x8c,
	0xd1, 0x0b, 0x82, 0xcb, 0x04, 0x21, 0xa9, 0x12, 0x92, 0xe5, 0xe2, 0xca, 0x07, 0x9b, 0x13, 0x9f,
	0x9b, 0x58, 0x20, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x13, 0xc4, 0x99, 0x0f, 0x33, 0xd9, 0x8a, 0xed,
	0xd3, 0x74, 0x0d, 0x26, 0x0e, 0x46, 0x25, 0x69, 0x2e, 0x49, 0x2c, 0xd6, 0x16, 0x17, 0xe4, 0xe7,
	0x15, 0xa7, 0x1a, 0xe5, 0x71, 0xf1, 0x39, 0x43, 0x9c, 0x1e, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c,
	0x2a, 0x14, 0xc3, 0x25, 0x88, 0xa1, 0x5c, 0x48, 0x01, 0xe6, 0x14, 0x5c, 0x1e, 0x90, 0x52, 0xc4,
	0xa3, 0x02, 0x62, 0x97, 0x12, 0x83, 0x06, 0xa3, 0x93, 0x41, 0x14, 0x48, 0x5d, 0x4e, 0x62, 0x92,
	0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x84, 0xa9, 0x9b, 0x5f, 0x94, 0xae, 0x0f, 0xd1, 0xad, 0x0b, 0x0e,
	0x29, 0xfd, 0xf4, 0x7c, 0x28, 0xbf, 0x20, 0x29, 0x89, 0x0d, 0x2c, 0x64, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x10, 0x0a, 0xea, 0x78, 0x63, 0x01, 0x00, 0x00,
}
