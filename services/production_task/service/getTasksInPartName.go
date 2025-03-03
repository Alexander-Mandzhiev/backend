package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/production_task"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) GetTasksInPartName(ctx context.Context, request *production_task.RequestTaskParams) (*production_task.ProductsResponse, error) {
	if request.Count <= 0 {
		sl.Log.Error("Count cannot be negative")
		return nil, status.Errorf(codes.InvalidArgument, "count cannot be negative")
	}

	products, err := s.productionTaskProvider.TaskInPartName(ctx, request)
	if err != nil {
		sl.Log.Error("Failed getting tasks in part name", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed getting tasks in part name: %v", err)
	}

	return &production_task.ProductsResponse{Products: products}, nil
}
