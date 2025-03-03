package handle

import (
	"backend/protos/gen/go/products_sk"
	"context"
	"google.golang.org/grpc"
)

type ProductSkService interface {
	Create(ctx context.Context, request *products_sk.CreateProductRequest) (*products_sk.ProductResponse, error)
	Product(ctx context.Context, request *products_sk.GetProductRequest) (*products_sk.ProductResponse, error)
	Update(ctx context.Context, request *products_sk.UpdateProductRequest) (*products_sk.ProductResponse, error)
	Delete(ctx context.Context, request *products_sk.DeleteProductRequest) (*products_sk.DeleteProductResponse, error)
	Products(ctx context.Context) (*products_sk.ProductListResponse, error)
	ProductsByStatus(ctx context.Context, statusID int32) (*products_sk.ProductListResponse, error)
}

type serverAPI struct {
	products_sk.UnimplementedProductServiceServer
	service ProductSkService
}

func Register(gRPCServer *grpc.Server, service ProductSkService) {
	products_sk.RegisterProductServiceServer(gRPCServer, &serverAPI{service: service})
}
