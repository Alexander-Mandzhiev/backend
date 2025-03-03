package handle

import (
	"backend/protos/gen/go/production_task"
	"context"
)

func (s *serverAPI) GetTasks(ctx context.Context, req *production_task.RequestTaskParams) (*production_task.ProductsResponse, error) {
	return s.service.GetTasks(ctx, req)
}
