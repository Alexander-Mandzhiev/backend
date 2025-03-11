package production_task_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/production_task"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *production_task.RequestTaskParams) (*production_task.ProductsResponse, error) {
	op := "ProductionTaskService.ListTasks"
	sl.Log.Debug("Calling gRPC ListTasks", slog.String("op", op))

	resp, err := s.client.GetTasks(ctx, req)
	if err != nil {
		sl.Log.Error("gRPC call failed", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	return resp, nil
}
