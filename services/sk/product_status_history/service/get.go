package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/product_status_history"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, request *product_status_history.GetStatusRequest) (*product_status_history.ProductSkStatusResponse, error) {
	op := "service.ProductSkStatus"
	sl.Log.Debug("Fetching product status", slog.String("op", op), slog.Int64("id", request.GetId()))

	status, err := s.provider.GetByID(ctx, request.GetId())
	if err != nil {
		if errors.Is(err, ErrProductStatusNotFound) {
			sl.Log.Warn("Product status not found", slog.Int64("id", request.GetId()), slog.String("op", op))
			return nil, err
		}
		sl.Log.Error("Failed to fetch product status", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Product status fetched successfully", slog.Int64("id", request.GetId()), slog.String("op", op))
	return status, nil
}
