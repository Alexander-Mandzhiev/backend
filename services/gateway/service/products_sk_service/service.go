package product_status_history

import (
	"backend/protos/gen/go/sk/products_sk"
	"google.golang.org/grpc"
)

type Service struct {
	client products_sk.ProductServiceClient
}

func New(conn *grpc.ClientConn) *Service {
	return &Service{client: products_sk.NewProductServiceClient(conn)}
}
