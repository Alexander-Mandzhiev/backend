package handle

import (
	"backend/protos/gen/go/production_task"
	"context"
)

func (s *serverAPI) GetTasksInPartName(ctx context.Context, req *production_task.RequestTaskParams) (*production_task.ProductsResponse, error) {
	return s.service.GetTasksInPartName(ctx, req)
}
