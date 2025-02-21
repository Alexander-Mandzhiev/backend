package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, request *statuses.CreateStatusRequest) (*statuses.StatusResponse, error) {
	op := "service.Create"
	sl.Log.Debug("Creating new status", slog.String("op", op))

	newStatus := &statuses.StatusResponse{
		Name:        request.Name,
		Description: request.Description,
	}

	id, err := s.statusesProvider.Create(ctx, newStatus)
	if err != nil {
		sl.Log.Error("Failed to create status", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	newStatus.Id = int32(id)
	sl.Log.Info("Status created successfully", slog.Int("id", id), slog.String("op", op))
	return newStatus, nil
}
