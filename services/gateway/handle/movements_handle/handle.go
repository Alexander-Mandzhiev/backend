package movements_handle

import (
	"backend/protos/gen/go/movements"
	"google.golang.org/grpc"
)

type MovementsHandle struct {
	client movements.MovementServiceClient
}

func New(conn *grpc.ClientConn) *MovementsHandle {
	return &MovementsHandle{client: movements.NewMovementServiceClient(conn)}
}
