package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/movements"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, request *movements.DeleteMovementRequest) (*movements.DeleteMovementResponse, error) {
	op := "service.Delete"
	sl.Log.Debug("Deleting movement", slog.Int64("id", request.Id), slog.String("op", op))
	if err := s.movementsProvider.Delete(ctx, request.Id); err != nil {
		if errors.Is(err, ErrMovementNotFound) {
			sl.Log.Warn("Movement not found for deletion", slog.Int64("id", request.Id), slog.String("op", op))
		} else {
			sl.Log.Error("Failed to delete movement", sl.Err(err), slog.Int64("id", request.Id), slog.String("op", op))
		}
		return nil, err
	}

	sl.Log.Info("Movement deleted successfully", slog.Int64("id", request.Id), slog.String("op", op))
	return &movements.DeleteMovementResponse{Success: true}, nil
}
