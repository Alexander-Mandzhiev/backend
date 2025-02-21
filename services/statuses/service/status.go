package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Status(ctx context.Context, request *statuses.GetStatusRequest) (*statuses.StatusResponse, error) {
	op := "service.Get"
	id := int(request.Id)
	sl.Log.Debug("Fetching status by ID", slog.Int("id", id), slog.String("op", op))

	status, err := s.statusesProvider.Status(ctx, id)
	if err != nil {
		if errors.Is(err, ErrStatusNotFound) {
			sl.Log.Warn("Status not found", slog.Int("id", id), slog.String("op", op))
		} else {
			sl.Log.Error("Failed to fetch status", sl.Err(err), slog.Int("id", id), slog.String("op", op))
		}
		return nil, err
	}

	sl.Log.Info("Status fetched successfully", slog.Int("id", int(status.Id)), slog.String("op", op))
	return status, nil
}
