// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	ListAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (User_ListAllUsersClient, error)
	NewUser(ctx context.Context, in *UserStruct, opts ...grpc.CallOption) (*Empty, error)
	RemoveUser(ctx context.Context, in *RemoveUserStruct, opts ...grpc.CallOption) (*Empty, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) ListAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (User_ListAllUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &User_ServiceDesc.Streams[0], "/user.User/ListAllUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userListAllUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type User_ListAllUsersClient interface {
	Recv() (*UserResponse, error)
	grpc.ClientStream
}

type userListAllUsersClient struct {
	grpc.ClientStream
}

func (x *userListAllUsersClient) Recv() (*UserResponse, error) {
	m := new(UserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userClient) NewUser(ctx context.Context, in *UserStruct, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/user.User/NewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RemoveUser(ctx context.Context, in *RemoveUserStruct, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/user.User/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	ListAllUsers(*Empty, User_ListAllUsersServer) error
	NewUser(context.Context, *UserStruct) (*Empty, error)
	RemoveUser(context.Context, *RemoveUserStruct) (*Empty, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) ListAllUsers(*Empty, User_ListAllUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAllUsers not implemented")
}
func (UnimplementedUserServer) NewUser(context.Context, *UserStruct) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewUser not implemented")
}
func (UnimplementedUserServer) RemoveUser(context.Context, *RemoveUserStruct) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_ListAllUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServer).ListAllUsers(m, &userListAllUsersServer{stream})
}

type User_ListAllUsersServer interface {
	Send(*UserResponse) error
	grpc.ServerStream
}

type userListAllUsersServer struct {
	grpc.ServerStream
}

func (x *userListAllUsersServer) Send(m *UserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _User_NewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserStruct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).NewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/NewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).NewUser(ctx, req.(*UserStruct))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserStruct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RemoveUser(ctx, req.(*RemoveUserStruct))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewUser",
			Handler:    _User_NewUser_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _User_RemoveUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAllUsers",
			Handler:       _User_ListAllUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/user/user.proto",
}
