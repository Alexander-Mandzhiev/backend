// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: production_task/production_task.proto

package production_task

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
	ProductionTaskService_GetTasks_FullMethodName           = "/production_task.ProductionTaskService/GetTasks"
	ProductionTaskService_GetTasksInPartName_FullMethodName = "/production_task.ProductionTaskService/GetTasksInPartName"
	ProductionTaskService_RecordInMsSQL_FullMethodName      = "/production_task.ProductionTaskService/RecordInMsSQL"
	ProductionTaskService_RecordOutMsSQL_FullMethodName     = "/production_task.ProductionTaskService/RecordOutMsSQL"
)

// ProductionTaskServiceClient is the client API for ProductionTaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductionTaskServiceClient interface {
	GetTasks(ctx context.Context, in *RequestTaskParams, opts ...grpc.CallOption) (*ProductsResponse, error)
	GetTasksInPartName(ctx context.Context, in *RequestTaskParams, opts ...grpc.CallOption) (*ProductsResponse, error)
	RecordInMsSQL(ctx context.Context, in *IDsRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	RecordOutMsSQL(ctx context.Context, in *IDsRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type productionTaskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductionTaskServiceClient(cc grpc.ClientConnInterface) ProductionTaskServiceClient {
	return &productionTaskServiceClient{cc}
}

func (c *productionTaskServiceClient) GetTasks(ctx context.Context, in *RequestTaskParams, opts ...grpc.CallOption) (*ProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductsResponse)
	err := c.cc.Invoke(ctx, ProductionTaskService_GetTasks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productionTaskServiceClient) GetTasksInPartName(ctx context.Context, in *RequestTaskParams, opts ...grpc.CallOption) (*ProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductsResponse)
	err := c.cc.Invoke(ctx, ProductionTaskService_GetTasksInPartName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productionTaskServiceClient) RecordInMsSQL(ctx context.Context, in *IDsRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, ProductionTaskService_RecordInMsSQL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productionTaskServiceClient) RecordOutMsSQL(ctx context.Context, in *IDsRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, ProductionTaskService_RecordOutMsSQL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductionTaskServiceServer is the server API for ProductionTaskService service.
// All implementations must embed UnimplementedProductionTaskServiceServer
// for forward compatibility.
type ProductionTaskServiceServer interface {
	GetTasks(context.Context, *RequestTaskParams) (*ProductsResponse, error)
	GetTasksInPartName(context.Context, *RequestTaskParams) (*ProductsResponse, error)
	RecordInMsSQL(context.Context, *IDsRequest) (*EmptyResponse, error)
	RecordOutMsSQL(context.Context, *IDsRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedProductionTaskServiceServer()
}

// UnimplementedProductionTaskServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductionTaskServiceServer struct{}

func (UnimplementedProductionTaskServiceServer) GetTasks(context.Context, *RequestTaskParams) (*ProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasks not implemented")
}
func (UnimplementedProductionTaskServiceServer) GetTasksInPartName(context.Context, *RequestTaskParams) (*ProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasksInPartName not implemented")
}
func (UnimplementedProductionTaskServiceServer) RecordInMsSQL(context.Context, *IDsRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordInMsSQL not implemented")
}
func (UnimplementedProductionTaskServiceServer) RecordOutMsSQL(context.Context, *IDsRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordOutMsSQL not implemented")
}
func (UnimplementedProductionTaskServiceServer) mustEmbedUnimplementedProductionTaskServiceServer() {}
func (UnimplementedProductionTaskServiceServer) testEmbeddedByValue()                               {}

// UnsafeProductionTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductionTaskServiceServer will
// result in compilation errors.
type UnsafeProductionTaskServiceServer interface {
	mustEmbedUnimplementedProductionTaskServiceServer()
}

func RegisterProductionTaskServiceServer(s grpc.ServiceRegistrar, srv ProductionTaskServiceServer) {
	// If the following call pancis, it indicates UnimplementedProductionTaskServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProductionTaskService_ServiceDesc, srv)
}

func _ProductionTaskService_GetTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTaskParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductionTaskServiceServer).GetTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductionTaskService_GetTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductionTaskServiceServer).GetTasks(ctx, req.(*RequestTaskParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductionTaskService_GetTasksInPartName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTaskParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductionTaskServiceServer).GetTasksInPartName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductionTaskService_GetTasksInPartName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductionTaskServiceServer).GetTasksInPartName(ctx, req.(*RequestTaskParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductionTaskService_RecordInMsSQL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductionTaskServiceServer).RecordInMsSQL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductionTaskService_RecordInMsSQL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductionTaskServiceServer).RecordInMsSQL(ctx, req.(*IDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductionTaskService_RecordOutMsSQL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductionTaskServiceServer).RecordOutMsSQL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductionTaskService_RecordOutMsSQL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductionTaskServiceServer).RecordOutMsSQL(ctx, req.(*IDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductionTaskService_ServiceDesc is the grpc.ServiceDesc for ProductionTaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductionTaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "production_task.ProductionTaskService",
	HandlerType: (*ProductionTaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTasks",
			Handler:    _ProductionTaskService_GetTasks_Handler,
		},
		{
			MethodName: "GetTasksInPartName",
			Handler:    _ProductionTaskService_GetTasksInPartName_Handler,
		},
		{
			MethodName: "RecordInMsSQL",
			Handler:    _ProductionTaskService_RecordInMsSQL_Handler,
		},
		{
			MethodName: "RecordOutMsSQL",
			Handler:    _ProductionTaskService_RecordOutMsSQL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "production_task/production_task.proto",
}
