package products_sk_statuses_service

import (
	"backend/protos/gen/go/products_sk_statuses"
	"google.golang.org/grpc"
)

type ProductsSkStatusesService struct {
	client products_sk_statuses.ProductStatusServiceClient
}

func New(conn *grpc.ClientConn) *ProductsSkStatusesService {
	return &ProductsSkStatusesService{client: products_sk_statuses.NewProductStatusServiceClient(conn)}
}
