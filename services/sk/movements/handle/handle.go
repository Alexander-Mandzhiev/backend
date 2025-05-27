package handle

import (
	"backend/protos/gen/go/sk/movements"
	"context"
	"google.golang.org/grpc"
)

type MovementsService interface {
	Create(ctx context.Context, request *movements.CreateMovementRequest) (*movements.MovementResponse, error)
	Movement(ctx context.Context, request *movements.GetMovementRequest) (*movements.MovementResponse, error)
	Update(ctx context.Context, request *movements.UpdateMovementRequest) (*movements.MovementResponse, error)
	Delete(ctx context.Context, request *movements.DeleteMovementRequest) (*movements.DeleteMovementResponse, error)
	Movements(ctx context.Context, request *movements.ListMovementsRequest) (*movements.MovementListResponse, error)
	InitializeProduct(ctx context.Context, request *movements.InitializeProductRequest) (*movements.InitializeProductResponse, error)
}
type serverAPI struct {
	movements.UnimplementedMovementServiceServer
	service MovementsService
}

func Register(gRPCServer *grpc.Server, service MovementsService) {
	movements.RegisterMovementServiceServer(gRPCServer, &serverAPI{service: service})
}
