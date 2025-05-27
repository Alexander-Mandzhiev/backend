package service

import (
	"context"
	"fmt"
	"log/slog"

	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/movements"
)

func (s *Service) InitializeProduct(ctx context.Context, req *movements.InitializeProductRequest) (*movements.InitializeProductResponse, error) {
	const op = "movements.Service.InitializeProduct"
	logger := sl.Log.With(slog.String("op", op))

	if _, err := s.locationProvider.Location(ctx, int(req.ToLocationId)); err != nil {
		return nil, fmt.Errorf("%s: invalid to_location_id %d: %w", op, req.ToLocationId, err)
	}
	if _, err := s.statusesProvider.Status(ctx, int(req.StatusId)); err != nil {
		return nil, fmt.Errorf("%s: invalid status_id %d: %w", op, req.StatusId, err)
	}

	response, err := s.movementsProvider.InitializeProduct(ctx, req)
	if err != nil {
		logger.Error("Failed to initialize products", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("Batch product initialization completed", slog.Int("created_count", len(response.CreatedProducts)), slog.Int("error_count", len(response.Errors)))
	return response, nil
}
