package products_sk_statuses_service

import (
	"backend/protos/gen/go/products_sk_statuses"
	"google.golang.org/grpc"
)

type Service struct {
	client products_sk_statuses.ProductStatusServiceClient
}

func New(conn *grpc.ClientConn) *Service {
	return &Service{client: products_sk_statuses.NewProductStatusServiceClient(conn)}
}
