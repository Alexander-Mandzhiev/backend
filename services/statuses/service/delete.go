package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, request *statuses.DeleteStatusRequest) (*statuses.DeleteStatusResponse, error) {
	op := "service.Delete"
	id := int(request.Id)
	sl.Log.Debug("Deleting status", slog.Int("id", id), slog.String("op", op))

	if err := s.statusesProvider.Delete(ctx, id); err != nil {
		if errors.Is(err, ErrStatusNotFound) {
			sl.Log.Warn("Status not found for deletion", slog.Int("id", id), slog.String("op", op))
		} else {
			sl.Log.Error("Failed to delete status", sl.Err(err), slog.Int("id", id), slog.String("op", op))
		}
		return nil, err
	}

	sl.Log.Info("Status deleted successfully", slog.Int("id", id), slog.String("op", op))
	return &statuses.DeleteStatusResponse{Success: true}, nil
}
