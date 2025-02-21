package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context) (*statuses.StatusListResponse, error) {
	op := "service.List"
	sl.Log.Debug("Fetching all statuses", slog.String("op", op))

	statusesList, err := s.statusesProvider.Statuses(ctx)
	if err != nil {
		sl.Log.Error("Failed to fetch statuses", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Statuses fetched successfully", slog.Int("count", len(statusesList)), slog.String("op", op))
	return &statuses.StatusListResponse{Data: statusesList}, nil
}
