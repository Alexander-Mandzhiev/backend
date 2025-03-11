package handle

import (
	"backend/protos/gen/go/production_task"
	"context"
	"google.golang.org/grpc"
)

type ProductionTaskService interface {
	GetTasks(ctx context.Context, request *production_task.RequestTaskParams) (*production_task.ProductsResponse, error)
	GetTasksInPartName(ctx context.Context, request *production_task.RequestTaskParams) (*production_task.ProductsResponse, error)
	RecordInMsSQL(ctx context.Context, ids []int64) error
	RecordOutMsSQL(ctx context.Context, ids []int64) error
}

type serverAPI struct {
	production_task.UnimplementedProductionTaskServiceServer
	service ProductionTaskService
}

func Register(gRPCServer *grpc.Server, service ProductionTaskService) {
	production_task.RegisterProductionTaskServiceServer(gRPCServer, &serverAPI{service: service})
}
