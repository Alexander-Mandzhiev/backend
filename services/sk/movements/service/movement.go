package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/movements"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Movement(ctx context.Context, request *movements.GetMovementRequest) (*movements.MovementResponse, error) {
	op := "service.Movement"
	sl.Log.Debug("Fetching movement by ID", slog.Int64("id", request.Id), slog.String("op", op))
	movement, err := s.movementsProvider.Movement(ctx, request.Id)
	if err != nil {
		if errors.Is(err, ErrMovementNotFound) {
			sl.Log.Warn("Movement not found", slog.Int64("id", request.Id), slog.String("op", op))
		} else {
			sl.Log.Error("Failed to fetch movement", sl.Err(err), slog.Int64("id", request.Id), slog.String("op", op))
		}
		return nil, err
	}

	sl.Log.Info("Movement fetched successfully", slog.Int64("id", movement.Id), slog.String("op", op))
	return movement, nil
}
