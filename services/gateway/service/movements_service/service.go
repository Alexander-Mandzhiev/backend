package movements_service

import (
	"backend/protos/gen/go/movements"
	"google.golang.org/grpc"
)

type Service struct {
	client movements.MovementServiceClient
}

func New(conn *grpc.ClientConn) *Service {
	return &Service{client: movements.NewMovementServiceClient(conn)}
}
