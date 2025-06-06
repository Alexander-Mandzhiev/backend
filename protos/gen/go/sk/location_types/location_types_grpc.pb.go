// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: sk/location_types/location_types.proto

package location_types

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
	LocationTypeService_CreateLocationType_FullMethodName = "/sk.location_types.LocationTypeService/CreateLocationType"
	LocationTypeService_GetLocationType_FullMethodName    = "/sk.location_types.LocationTypeService/GetLocationType"
	LocationTypeService_UpdateLocationType_FullMethodName = "/sk.location_types.LocationTypeService/UpdateLocationType"
	LocationTypeService_DeleteLocationType_FullMethodName = "/sk.location_types.LocationTypeService/DeleteLocationType"
	LocationTypeService_ListLocationType_FullMethodName   = "/sk.location_types.LocationTypeService/ListLocationType"
)

// LocationTypeServiceClient is the client API for LocationTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationTypeServiceClient interface {
	CreateLocationType(ctx context.Context, in *CreateLocationTypeRequest, opts ...grpc.CallOption) (*LocationTypeResponse, error)
	GetLocationType(ctx context.Context, in *GetLocationTypeRequest, opts ...grpc.CallOption) (*LocationTypeResponse, error)
	UpdateLocationType(ctx context.Context, in *UpdateLocationTypeRequest, opts ...grpc.CallOption) (*LocationTypeResponse, error)
	DeleteLocationType(ctx context.Context, in *DeleteLocationTypeRequest, opts ...grpc.CallOption) (*DeleteLocationTypeResponse, error)
	ListLocationType(ctx context.Context, in *ListLocationTypesRequest, opts ...grpc.CallOption) (*LocationTypeListResponse, error)
}

type locationTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationTypeServiceClient(cc grpc.ClientConnInterface) LocationTypeServiceClient {
	return &locationTypeServiceClient{cc}
}

func (c *locationTypeServiceClient) CreateLocationType(ctx context.Context, in *CreateLocationTypeRequest, opts ...grpc.CallOption) (*LocationTypeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LocationTypeResponse)
	err := c.cc.Invoke(ctx, LocationTypeService_CreateLocationType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationTypeServiceClient) GetLocationType(ctx context.Context, in *GetLocationTypeRequest, opts ...grpc.CallOption) (*LocationTypeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LocationTypeResponse)
	err := c.cc.Invoke(ctx, LocationTypeService_GetLocationType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationTypeServiceClient) UpdateLocationType(ctx context.Context, in *UpdateLocationTypeRequest, opts ...grpc.CallOption) (*LocationTypeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LocationTypeResponse)
	err := c.cc.Invoke(ctx, LocationTypeService_UpdateLocationType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationTypeServiceClient) DeleteLocationType(ctx context.Context, in *DeleteLocationTypeRequest, opts ...grpc.CallOption) (*DeleteLocationTypeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteLocationTypeResponse)
	err := c.cc.Invoke(ctx, LocationTypeService_DeleteLocationType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationTypeServiceClient) ListLocationType(ctx context.Context, in *ListLocationTypesRequest, opts ...grpc.CallOption) (*LocationTypeListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LocationTypeListResponse)
	err := c.cc.Invoke(ctx, LocationTypeService_ListLocationType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationTypeServiceServer is the server API for LocationTypeService service.
// All implementations must embed UnimplementedLocationTypeServiceServer
// for forward compatibility.
type LocationTypeServiceServer interface {
	CreateLocationType(context.Context, *CreateLocationTypeRequest) (*LocationTypeResponse, error)
	GetLocationType(context.Context, *GetLocationTypeRequest) (*LocationTypeResponse, error)
	UpdateLocationType(context.Context, *UpdateLocationTypeRequest) (*LocationTypeResponse, error)
	DeleteLocationType(context.Context, *DeleteLocationTypeRequest) (*DeleteLocationTypeResponse, error)
	ListLocationType(context.Context, *ListLocationTypesRequest) (*LocationTypeListResponse, error)
	mustEmbedUnimplementedLocationTypeServiceServer()
}

// UnimplementedLocationTypeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLocationTypeServiceServer struct{}

func (UnimplementedLocationTypeServiceServer) CreateLocationType(context.Context, *CreateLocationTypeRequest) (*LocationTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLocationType not implemented")
}
func (UnimplementedLocationTypeServiceServer) GetLocationType(context.Context, *GetLocationTypeRequest) (*LocationTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocationType not implemented")
}
func (UnimplementedLocationTypeServiceServer) UpdateLocationType(context.Context, *UpdateLocationTypeRequest) (*LocationTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLocationType not implemented")
}
func (UnimplementedLocationTypeServiceServer) DeleteLocationType(context.Context, *DeleteLocationTypeRequest) (*DeleteLocationTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocationType not implemented")
}
func (UnimplementedLocationTypeServiceServer) ListLocationType(context.Context, *ListLocationTypesRequest) (*LocationTypeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLocationType not implemented")
}
func (UnimplementedLocationTypeServiceServer) mustEmbedUnimplementedLocationTypeServiceServer() {}
func (UnimplementedLocationTypeServiceServer) testEmbeddedByValue()                             {}

// UnsafeLocationTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationTypeServiceServer will
// result in compilation errors.
type UnsafeLocationTypeServiceServer interface {
	mustEmbedUnimplementedLocationTypeServiceServer()
}

func RegisterLocationTypeServiceServer(s grpc.ServiceRegistrar, srv LocationTypeServiceServer) {
	// If the following call pancis, it indicates UnimplementedLocationTypeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LocationTypeService_ServiceDesc, srv)
}

func _LocationTypeService_CreateLocationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLocationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationTypeServiceServer).CreateLocationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationTypeService_CreateLocationType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationTypeServiceServer).CreateLocationType(ctx, req.(*CreateLocationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationTypeService_GetLocationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationTypeServiceServer).GetLocationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationTypeService_GetLocationType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationTypeServiceServer).GetLocationType(ctx, req.(*GetLocationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationTypeService_UpdateLocationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLocationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationTypeServiceServer).UpdateLocationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationTypeService_UpdateLocationType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationTypeServiceServer).UpdateLocationType(ctx, req.(*UpdateLocationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationTypeService_DeleteLocationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLocationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationTypeServiceServer).DeleteLocationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationTypeService_DeleteLocationType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationTypeServiceServer).DeleteLocationType(ctx, req.(*DeleteLocationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationTypeService_ListLocationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLocationTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationTypeServiceServer).ListLocationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationTypeService_ListLocationType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationTypeServiceServer).ListLocationType(ctx, req.(*ListLocationTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocationTypeService_ServiceDesc is the grpc.ServiceDesc for LocationTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sk.location_types.LocationTypeService",
	HandlerType: (*LocationTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLocationType",
			Handler:    _LocationTypeService_CreateLocationType_Handler,
		},
		{
			MethodName: "GetLocationType",
			Handler:    _LocationTypeService_GetLocationType_Handler,
		},
		{
			MethodName: "UpdateLocationType",
			Handler:    _LocationTypeService_UpdateLocationType_Handler,
		},
		{
			MethodName: "DeleteLocationType",
			Handler:    _LocationTypeService_DeleteLocationType_Handler,
		},
		{
			MethodName: "ListLocationType",
			Handler:    _LocationTypeService_ListLocationType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sk/location_types/location_types.proto",
}
