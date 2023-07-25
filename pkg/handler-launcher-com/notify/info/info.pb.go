// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/handler-launcher-com/notify/info/info.proto

/*
Package info is a generated protocol buffer package.

It is generated from these files:

	pkg/handler-launcher-com/notify/info/info.proto

It has these top-level messages:

	NotifyInfoRequest
	NotifyInfoResponse
*/
package info

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

type NotifyInfoRequest struct {
}

func (m *NotifyInfoRequest) Reset()                    { *m = NotifyInfoRequest{} }
func (m *NotifyInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*NotifyInfoRequest) ProtoMessage()               {}
func (*NotifyInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type NotifyInfoResponse struct {
	SupportedNotifyVersions []uint32 `protobuf:"varint,1,rep,packed,name=supportedNotifyVersions" json:"supportedNotifyVersions,omitempty"`
}

func (m *NotifyInfoResponse) Reset()                    { *m = NotifyInfoResponse{} }
func (m *NotifyInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*NotifyInfoResponse) ProtoMessage()               {}
func (*NotifyInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NotifyInfoResponse) GetSupportedNotifyVersions() []uint32 {
	if m != nil {
		return m.SupportedNotifyVersions
	}
	return nil
}

func init() {
	proto.RegisterType((*NotifyInfoRequest)(nil), "kubevirt.notify.info.NotifyInfoRequest")
	proto.RegisterType((*NotifyInfoResponse)(nil), "kubevirt.notify.info.NotifyInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NotifyInfo service

type NotifyInfoClient interface {
	Info(ctx context.Context, in *NotifyInfoRequest, opts ...grpc.CallOption) (*NotifyInfoResponse, error)
}

type notifyInfoClient struct {
	cc *grpc.ClientConn
}

func NewNotifyInfoClient(cc *grpc.ClientConn) NotifyInfoClient {
	return &notifyInfoClient{cc}
}

func (c *notifyInfoClient) Info(ctx context.Context, in *NotifyInfoRequest, opts ...grpc.CallOption) (*NotifyInfoResponse, error) {
	out := new(NotifyInfoResponse)
	err := grpc.Invoke(ctx, "/kubevirt.notify.info.NotifyInfo/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NotifyInfo service

type NotifyInfoServer interface {
	Info(context.Context, *NotifyInfoRequest) (*NotifyInfoResponse, error)
}

func RegisterNotifyInfoServer(s *grpc.Server, srv NotifyInfoServer) {
	s.RegisterService(&_NotifyInfo_serviceDesc, srv)
}

func _NotifyInfo_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyInfoServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.notify.info.NotifyInfo/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyInfoServer).Info(ctx, req.(*NotifyInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotifyInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kubevirt.notify.info.NotifyInfo",
	HandlerType: (*NotifyInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _NotifyInfo_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/handler-launcher-com/notify/info/info.proto",
}

func init() { proto.RegisterFile("pkg/handler-launcher-com/notify/info/info.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2f, 0xc8, 0x4e, 0xd7,
	0xcf, 0x48, 0xcc, 0x4b, 0xc9, 0x49, 0x2d, 0xd2, 0xcd, 0x49, 0x2c, 0xcd, 0x4b, 0xce, 0x48, 0x2d,
	0xd2, 0x4d, 0xce, 0xcf, 0xd5, 0xcf, 0xcb, 0x2f, 0xc9, 0x4c, 0xab, 0xd4, 0xcf, 0xcc, 0x4b, 0xcb,
	0x07, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x22, 0xd9, 0xa5, 0x49, 0xa9, 0x65, 0x99,
	0x45, 0x25, 0x7a, 0x10, 0x05, 0x7a, 0x20, 0x39, 0x25, 0x61, 0x2e, 0x41, 0x3f, 0x30, 0xd7, 0x33,
	0x2f, 0x2d, 0x3f, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0xc9, 0x8f, 0x4b, 0x08, 0x59, 0xb0,
	0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc8, 0x82, 0x4b, 0xbc, 0xb8, 0xb4, 0xa0, 0x20, 0xbf, 0xa8,
	0x24, 0x35, 0x05, 0x22, 0x1d, 0x96, 0x5a, 0x54, 0x9c, 0x99, 0x9f, 0x57, 0x2c, 0xc1, 0xa8, 0xc0,
	0xac, 0xc1, 0x1b, 0x84, 0x4b, 0xda, 0x28, 0x93, 0x8b, 0x0b, 0x61, 0x9e, 0x50, 0x34, 0x17, 0x0b,
	0x98, 0x56, 0xd7, 0xc3, 0xe6, 0x22, 0x3d, 0x0c, 0xe7, 0x48, 0x69, 0x10, 0x56, 0x08, 0x71, 0xa2,
	0x12, 0x83, 0x13, 0x5b, 0x14, 0x0b, 0x48, 0x32, 0x89, 0x0d, 0xec, 0x69, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x33, 0x3f, 0x77, 0xf0, 0x27, 0x01, 0x00, 0x00,
}
