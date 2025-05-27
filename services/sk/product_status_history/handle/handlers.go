package handle

import (
	"backend/protos/gen/go/sk/product_status_history"
	"context"
	"google.golang.org/grpc"
)

type StatusSkStatusesService interface {
	Create(ctx context.Context, request *product_status_history.CreateStatusRequest) (*product_status_history.ProductSkStatusResponse, error)
	Get(ctx context.Context, request *product_status_history.GetStatusRequest) (*product_status_history.ProductSkStatusResponse, error)
	Delete(ctx context.Context, request *product_status_history.DeleteStatusRequest) (*product_status_history.DeleteStatusResponse, error)
	List(ctx context.Context, request *product_status_history.ListStatusesByProductRequest) (*product_status_history.StatusListResponse, error)
}

type serverAPI struct {
	product_status_history.UnimplementedProductStatusServiceServer
	service StatusSkStatusesService
}

func Register(gRPCServer *grpc.Server, service StatusSkStatusesService) {
	product_status_history.RegisterProductStatusServiceServer(gRPCServer, &serverAPI{service: service})
}
