package production_task_handle

import (
	"backend/protos/gen/go/production_task"
	"google.golang.org/grpc"
)

type ProductionTaskHandle struct {
	client production_task.ProductionTaskServiceClient
}

func New(conn *grpc.ClientConn) *ProductionTaskHandle {
	return &ProductionTaskHandle{client: production_task.NewProductionTaskServiceClient(conn)}
}
