package products_sk_service

import (
	"backend/protos/gen/go/products_sk"
	"google.golang.org/grpc"
)

type Service struct {
	client products_sk.ProductServiceClient
}

func New(conn *grpc.ClientConn) *Service {
	return &Service{client: products_sk.NewProductServiceClient(conn)}
}
