package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/movements"
	"context"
	"log/slog"
)

func (s *Service) Movements(ctx context.Context, request *movements.ListMovementsRequest) (*movements.MovementListResponse, error) {
	op := "service.List"
	sl.Log.Debug("Fetching all movements by product ID", slog.Int64("product Id", request.ProductId), slog.String("op", op))

	movementsList, err := s.movementsProvider.Movements(ctx, request.ProductId)
	if err != nil {
		sl.Log.Error("Failed to fetch movements", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Movements fetched successfully", slog.Int("count", len(movementsList)), slog.String("op", op))
	return &movements.MovementListResponse{Data: movementsList}, nil
}
