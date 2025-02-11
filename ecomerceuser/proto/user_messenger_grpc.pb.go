// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/user_messenger.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserMessenger_GetUserList_FullMethodName = "/UserMessenger/GetUserList"
)

// UserMessengerClient is the client API for UserMessenger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserMessengerClient interface {
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error)
}

type userMessengerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserMessengerClient(cc grpc.ClientConnInterface) UserMessengerClient {
	return &userMessengerClient{cc}
}

func (c *userMessengerClient) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserListResponse)
	err := c.cc.Invoke(ctx, UserMessenger_GetUserList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserMessengerServer is the server API for UserMessenger service.
// All implementations must embed UnimplementedUserMessengerServer
// for forward compatibility.
type UserMessengerServer interface {
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error)
	mustEmbedUnimplementedUserMessengerServer()
}

// UnimplementedUserMessengerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserMessengerServer struct{}

func (UnimplementedUserMessengerServer) GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedUserMessengerServer) mustEmbedUnimplementedUserMessengerServer() {}
func (UnimplementedUserMessengerServer) testEmbeddedByValue()                       {}

// UnsafeUserMessengerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserMessengerServer will
// result in compilation errors.
type UnsafeUserMessengerServer interface {
	mustEmbedUnimplementedUserMessengerServer()
}

func RegisterUserMessengerServer(s grpc.ServiceRegistrar, srv UserMessengerServer) {
	// If the following call pancis, it indicates UnimplementedUserMessengerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserMessenger_ServiceDesc, srv)
}

func _UserMessenger_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMessengerServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserMessenger_GetUserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMessengerServer).GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserMessenger_ServiceDesc is the grpc.ServiceDesc for UserMessenger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserMessenger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserMessenger",
	HandlerType: (*UserMessengerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _UserMessenger_GetUserList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user_messenger.proto",
}
