package product_status_history

import (
	"backend/protos/gen/go/sk/product_status_history"
	"google.golang.org/grpc"
)

type Service struct {
	client product_status_history.ProductStatusServiceClient
}

func New(conn *grpc.ClientConn) *Service {
	return &Service{client: product_status_history.NewProductStatusServiceClient(conn)}
}
