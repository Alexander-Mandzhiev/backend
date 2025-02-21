package handle

import (
	"backend/protos/gen/go/statuses"
	"context"
	"google.golang.org/grpc"
)

type StatusesService interface {
	Create(ctx context.Context, request *statuses.CreateStatusRequest) (*statuses.StatusResponse, error)
	Status(ctx context.Context, request *statuses.GetStatusRequest) (*statuses.StatusResponse, error)
	Update(ctx context.Context, request *statuses.UpdateStatusRequest) (*statuses.StatusResponse, error)
	Delete(ctx context.Context, request *statuses.DeleteStatusRequest) (*statuses.DeleteStatusResponse, error)
	List(ctx context.Context) (*statuses.StatusListResponse, error)
}
type serverAPI struct {
	statuses.UnimplementedStatusServiceServer
	service StatusesService
}

func Register(gRPCServer *grpc.Server, service StatusesService) {
	statuses.RegisterStatusServiceServer(gRPCServer, &serverAPI{service: service})
}
