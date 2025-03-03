package handle

import (
	"backend/protos/gen/go/products_sk_statuses"
	"context"
	"google.golang.org/grpc"
)

type StatusSkStatusesService interface {
	Create(ctx context.Context, request *products_sk_statuses.CreateProductStatusRequest) (*products_sk_statuses.ProductSkStatusResponse, error)
	Get(ctx context.Context, request *products_sk_statuses.GetProductStatusRequest) (*products_sk_statuses.ProductSkStatusResponse, error)
	Update(ctx context.Context, request *products_sk_statuses.UpdateProductStatusRequest) (*products_sk_statuses.ProductSkStatusResponse, error)
	Delete(ctx context.Context, request *products_sk_statuses.DeleteProductStatusRequest) (*products_sk_statuses.DeleteProductStatusResponse, error)
	List(ctx context.Context, request *products_sk_statuses.ListProductStatusesRequest) (*products_sk_statuses.ProductStatusListResponse, error)
}

type serverAPI struct {
	products_sk_statuses.UnimplementedProductStatusServiceServer
	service StatusSkStatusesService
}

func Register(gRPCServer *grpc.Server, service StatusSkStatusesService) {
	products_sk_statuses.RegisterProductStatusServiceServer(gRPCServer, &serverAPI{service: service})
}
