// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: sk/statuses/statuses.proto

package statuses

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
	StatusService_CreateStatus_FullMethodName = "/sk.statuses.StatusService/CreateStatus"
	StatusService_GetStatus_FullMethodName    = "/sk.statuses.StatusService/GetStatus"
	StatusService_UpdateStatus_FullMethodName = "/sk.statuses.StatusService/UpdateStatus"
	StatusService_DeleteStatus_FullMethodName = "/sk.statuses.StatusService/DeleteStatus"
	StatusService_ListStatuses_FullMethodName = "/sk.statuses.StatusService/ListStatuses"
)

// StatusServiceClient is the client API for StatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// CRUD методы для статусов
type StatusServiceClient interface {
	CreateStatus(ctx context.Context, in *CreateStatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	DeleteStatus(ctx context.Context, in *DeleteStatusRequest, opts ...grpc.CallOption) (*DeleteStatusResponse, error)
	ListStatuses(ctx context.Context, in *ListStatusesRequest, opts ...grpc.CallOption) (*StatusListResponse, error)
}

type statusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatusServiceClient(cc grpc.ClientConnInterface) StatusServiceClient {
	return &statusServiceClient{cc}
}

func (c *statusServiceClient) CreateStatus(ctx context.Context, in *CreateStatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, StatusService_CreateStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusServiceClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, StatusService_GetStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusServiceClient) UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, StatusService_UpdateStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusServiceClient) DeleteStatus(ctx context.Context, in *DeleteStatusRequest, opts ...grpc.CallOption) (*DeleteStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteStatusResponse)
	err := c.cc.Invoke(ctx, StatusService_DeleteStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusServiceClient) ListStatuses(ctx context.Context, in *ListStatusesRequest, opts ...grpc.CallOption) (*StatusListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusListResponse)
	err := c.cc.Invoke(ctx, StatusService_ListStatuses_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatusServiceServer is the server API for StatusService service.
// All implementations must embed UnimplementedStatusServiceServer
// for forward compatibility.
//
// CRUD методы для статусов
type StatusServiceServer interface {
	CreateStatus(context.Context, *CreateStatusRequest) (*StatusResponse, error)
	GetStatus(context.Context, *GetStatusRequest) (*StatusResponse, error)
	UpdateStatus(context.Context, *UpdateStatusRequest) (*StatusResponse, error)
	DeleteStatus(context.Context, *DeleteStatusRequest) (*DeleteStatusResponse, error)
	ListStatuses(context.Context, *ListStatusesRequest) (*StatusListResponse, error)
	mustEmbedUnimplementedStatusServiceServer()
}

// UnimplementedStatusServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStatusServiceServer struct{}

func (UnimplementedStatusServiceServer) CreateStatus(context.Context, *CreateStatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStatus not implemented")
}
func (UnimplementedStatusServiceServer) GetStatus(context.Context, *GetStatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedStatusServiceServer) UpdateStatus(context.Context, *UpdateStatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedStatusServiceServer) DeleteStatus(context.Context, *DeleteStatusRequest) (*DeleteStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStatus not implemented")
}
func (UnimplementedStatusServiceServer) ListStatuses(context.Context, *ListStatusesRequest) (*StatusListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStatuses not implemented")
}
func (UnimplementedStatusServiceServer) mustEmbedUnimplementedStatusServiceServer() {}
func (UnimplementedStatusServiceServer) testEmbeddedByValue()                       {}

// UnsafeStatusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatusServiceServer will
// result in compilation errors.
type UnsafeStatusServiceServer interface {
	mustEmbedUnimplementedStatusServiceServer()
}

func RegisterStatusServiceServer(s grpc.ServiceRegistrar, srv StatusServiceServer) {
	// If the following call pancis, it indicates UnimplementedStatusServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StatusService_ServiceDesc, srv)
}

func _StatusService_CreateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServiceServer).CreateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatusService_CreateStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServiceServer).CreateStatus(ctx, req.(*CreateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatusService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatusService_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServiceServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatusService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatusService_UpdateStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServiceServer).UpdateStatus(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatusService_DeleteStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServiceServer).DeleteStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatusService_DeleteStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServiceServer).DeleteStatus(ctx, req.(*DeleteStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatusService_ListStatuses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStatusesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServiceServer).ListStatuses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatusService_ListStatuses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServiceServer).ListStatuses(ctx, req.(*ListStatusesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatusService_ServiceDesc is the grpc.ServiceDesc for StatusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sk.statuses.StatusService",
	HandlerType: (*StatusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStatus",
			Handler:    _StatusService_CreateStatus_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _StatusService_GetStatus_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _StatusService_UpdateStatus_Handler,
		},
		{
			MethodName: "DeleteStatus",
			Handler:    _StatusService_DeleteStatus_Handler,
		},
		{
			MethodName: "ListStatuses",
			Handler:    _StatusService_ListStatuses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sk/statuses/statuses.proto",
}
