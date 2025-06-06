// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: apps/apps.proto

package app_provider

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
	AppProviderService_Create_FullMethodName = "/app_provider.AppProviderService/Create"
	AppProviderService_App_FullMethodName    = "/app_provider.AppProviderService/App"
	AppProviderService_Apps_FullMethodName   = "/app_provider.AppProviderService/Apps"
	AppProviderService_Update_FullMethodName = "/app_provider.AppProviderService/Update"
	AppProviderService_Delete_FullMethodName = "/app_provider.AppProviderService/Delete"
)

// AppProviderServiceClient is the client API for AppProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppProviderServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	App(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*GetAppResponse, error)
	Apps(ctx context.Context, in *GetAppsRequest, opts ...grpc.CallOption) (*GetAppsResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type appProviderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppProviderServiceClient(cc grpc.ClientConnInterface) AppProviderServiceClient {
	return &appProviderServiceClient{cc}
}

func (c *appProviderServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, AppProviderService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appProviderServiceClient) App(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*GetAppResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAppResponse)
	err := c.cc.Invoke(ctx, AppProviderService_App_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appProviderServiceClient) Apps(ctx context.Context, in *GetAppsRequest, opts ...grpc.CallOption) (*GetAppsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAppsResponse)
	err := c.cc.Invoke(ctx, AppProviderService_Apps_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appProviderServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, AppProviderService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appProviderServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, AppProviderService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppProviderServiceServer is the server API for AppProviderService service.
// All implementations must embed UnimplementedAppProviderServiceServer
// for forward compatibility.
type AppProviderServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	App(context.Context, *GetAppRequest) (*GetAppResponse, error)
	Apps(context.Context, *GetAppsRequest) (*GetAppsResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedAppProviderServiceServer()
}

// UnimplementedAppProviderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAppProviderServiceServer struct{}

func (UnimplementedAppProviderServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAppProviderServiceServer) App(context.Context, *GetAppRequest) (*GetAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method App not implemented")
}
func (UnimplementedAppProviderServiceServer) Apps(context.Context, *GetAppsRequest) (*GetAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apps not implemented")
}
func (UnimplementedAppProviderServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAppProviderServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAppProviderServiceServer) mustEmbedUnimplementedAppProviderServiceServer() {}
func (UnimplementedAppProviderServiceServer) testEmbeddedByValue()                            {}

// UnsafeAppProviderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppProviderServiceServer will
// result in compilation errors.
type UnsafeAppProviderServiceServer interface {
	mustEmbedUnimplementedAppProviderServiceServer()
}

func RegisterAppProviderServiceServer(s grpc.ServiceRegistrar, srv AppProviderServiceServer) {
	// If the following call pancis, it indicates UnimplementedAppProviderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AppProviderService_ServiceDesc, srv)
}

func _AppProviderService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppProviderServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppProviderService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppProviderServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppProviderService_App_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppProviderServiceServer).App(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppProviderService_App_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppProviderServiceServer).App(ctx, req.(*GetAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppProviderService_Apps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppProviderServiceServer).Apps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppProviderService_Apps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppProviderServiceServer).Apps(ctx, req.(*GetAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppProviderService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppProviderServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppProviderService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppProviderServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppProviderService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppProviderServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppProviderService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppProviderServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppProviderService_ServiceDesc is the grpc.ServiceDesc for AppProviderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppProviderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app_provider.AppProviderService",
	HandlerType: (*AppProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AppProviderService_Create_Handler,
		},
		{
			MethodName: "App",
			Handler:    _AppProviderService_App_Handler,
		},
		{
			MethodName: "Apps",
			Handler:    _AppProviderService_Apps_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AppProviderService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AppProviderService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/apps.proto",
}
