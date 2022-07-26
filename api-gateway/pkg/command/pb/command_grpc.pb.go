// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api-gateway/pkg/command/pb/command.proto

package pb

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

// CommandServiceClient is the client API for CommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandServiceClient interface {
	//rpc Command(CommandRequest) returns (CommandResponse) {}
	AddCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error)
	ModifyCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error)
	DeleteCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error)
}

type commandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandServiceClient(cc grpc.ClientConnInterface) CommandServiceClient {
	return &commandServiceClient{cc}
}

func (c *commandServiceClient) AddCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, "/command.CommandService/AddCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandServiceClient) ModifyCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, "/command.CommandService/ModifyCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandServiceClient) DeleteCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, "/command.CommandService/DeleteCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandServiceServer is the server API for CommandService service.
// All implementations must embed UnimplementedCommandServiceServer
// for forward compatibility
type CommandServiceServer interface {
	//rpc Command(CommandRequest) returns (CommandResponse) {}
	AddCommand(context.Context, *CommandRequest) (*CommandResponse, error)
	ModifyCommand(context.Context, *CommandRequest) (*CommandResponse, error)
	DeleteCommand(context.Context, *CommandRequest) (*CommandResponse, error)
	mustEmbedUnimplementedCommandServiceServer()
}

// UnimplementedCommandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommandServiceServer struct {
}

func (UnimplementedCommandServiceServer) AddCommand(context.Context, *CommandRequest) (*CommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCommand not implemented")
}
func (UnimplementedCommandServiceServer) ModifyCommand(context.Context, *CommandRequest) (*CommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyCommand not implemented")
}
func (UnimplementedCommandServiceServer) DeleteCommand(context.Context, *CommandRequest) (*CommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommand not implemented")
}
func (UnimplementedCommandServiceServer) mustEmbedUnimplementedCommandServiceServer() {}

// UnsafeCommandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandServiceServer will
// result in compilation errors.
type UnsafeCommandServiceServer interface {
	mustEmbedUnimplementedCommandServiceServer()
}

func RegisterCommandServiceServer(s grpc.ServiceRegistrar, srv CommandServiceServer) {
	s.RegisterService(&CommandService_ServiceDesc, srv)
}

func _CommandService_AddCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).AddCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/command.CommandService/AddCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).AddCommand(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandService_ModifyCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).ModifyCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/command.CommandService/ModifyCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).ModifyCommand(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandService_DeleteCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).DeleteCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/command.CommandService/DeleteCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).DeleteCommand(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommandService_ServiceDesc is the grpc.ServiceDesc for CommandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "command.CommandService",
	HandlerType: (*CommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCommand",
			Handler:    _CommandService_AddCommand_Handler,
		},
		{
			MethodName: "ModifyCommand",
			Handler:    _CommandService_ModifyCommand_Handler,
		},
		{
			MethodName: "DeleteCommand",
			Handler:    _CommandService_DeleteCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api-gateway/pkg/command/pb/command.proto",
}
