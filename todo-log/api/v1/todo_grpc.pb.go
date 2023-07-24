// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: todo-log/api/v1/todo.proto

package api_v1

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

const (
	Todo_Log_Insert_FullMethodName   = "/api.v1.Todo_Log/Insert"
	Todo_Log_Retrieve_FullMethodName = "/api.v1.Todo_Log/Retrieve"
	Todo_Log_List_FullMethodName     = "/api.v1.Todo_Log/List"
	Todo_Log_Update_FullMethodName   = "/api.v1.Todo_Log/Update"
)

// Todo_LogClient is the client API for Todo_Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Todo_LogClient interface {
	Insert(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Response, error)
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*Todo, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Activities, error)
	Update(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Activities, error)
}

type todo_LogClient struct {
	cc grpc.ClientConnInterface
}

func NewTodo_LogClient(cc grpc.ClientConnInterface) Todo_LogClient {
	return &todo_LogClient{cc}
}

func (c *todo_LogClient) Insert(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Todo_Log_Insert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todo_LogClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, Todo_Log_Retrieve_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todo_LogClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Activities, error) {
	out := new(Activities)
	err := c.cc.Invoke(ctx, Todo_Log_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todo_LogClient) Update(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Activities, error) {
	out := new(Activities)
	err := c.cc.Invoke(ctx, Todo_Log_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Todo_LogServer is the server API for Todo_Log service.
// All implementations must embed UnimplementedTodo_LogServer
// for forward compatibility
type Todo_LogServer interface {
	Insert(context.Context, *Todo) (*Response, error)
	Retrieve(context.Context, *RetrieveRequest) (*Todo, error)
	List(context.Context, *ListRequest) (*Activities, error)
	Update(context.Context, *Todo) (*Activities, error)
	mustEmbedUnimplementedTodo_LogServer()
}

// UnimplementedTodo_LogServer must be embedded to have forward compatible implementations.
type UnimplementedTodo_LogServer struct {
}

func (UnimplementedTodo_LogServer) Insert(context.Context, *Todo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedTodo_LogServer) Retrieve(context.Context, *RetrieveRequest) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedTodo_LogServer) List(context.Context, *ListRequest) (*Activities, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTodo_LogServer) Update(context.Context, *Todo) (*Activities, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTodo_LogServer) mustEmbedUnimplementedTodo_LogServer() {}

// UnsafeTodo_LogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Todo_LogServer will
// result in compilation errors.
type UnsafeTodo_LogServer interface {
	mustEmbedUnimplementedTodo_LogServer()
}

func RegisterTodo_LogServer(s grpc.ServiceRegistrar, srv Todo_LogServer) {
	s.RegisterService(&Todo_Log_ServiceDesc, srv)
}

func _Todo_Log_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Todo_LogServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Todo_Log_Insert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Todo_LogServer).Insert(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_Log_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Todo_LogServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Todo_Log_Retrieve_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Todo_LogServer).Retrieve(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_Log_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Todo_LogServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Todo_Log_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Todo_LogServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_Log_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Todo_LogServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Todo_Log_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Todo_LogServer).Update(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

// Todo_Log_ServiceDesc is the grpc.ServiceDesc for Todo_Log service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Todo_Log_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.Todo_Log",
	HandlerType: (*Todo_LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _Todo_Log_Insert_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _Todo_Log_Retrieve_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Todo_Log_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Todo_Log_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo-log/api/v1/todo.proto",
}
