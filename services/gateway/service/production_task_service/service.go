package production_task_service

import (
	"backend/protos/gen/go/production_task"
	"context"
	"google.golang.org/grpc"
)

type ProductionTaskService interface {
	GetTasks(ctx context.Context, req *production_task.RequestTaskParams) (*production_task.ProductsResponse, error)
	GetTasksByPartName(ctx context.Context, req *production_task.RequestTaskParams) (*production_task.ProductsResponse, error)
}

type Service struct {
	client production_task.ProductionTaskServiceClient
}

func New(conn *grpc.ClientConn) *Service {
	return &Service{client: production_task.NewProductionTaskServiceClient(conn)}
}
