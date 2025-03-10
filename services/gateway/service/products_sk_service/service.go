package products_sk_service

import (
	"backend/protos/gen/go/products_sk"
	"google.golang.org/grpc"
)

type ProductionSkHandle struct {
	client products_sk.ProductServiceClient
}

func New(conn *grpc.ClientConn) *ProductionSkHandle {
	return &ProductionSkHandle{client: products_sk.NewProductServiceClient(conn)}
}
