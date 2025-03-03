package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk_statuses"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, request *products_sk_statuses.GetProductStatusRequest) (*products_sk_statuses.ProductSkStatusResponse, error) {
	op := "service.ProductSkStatus"
	sl.Log.Debug("Fetching product status", slog.String("op", op), slog.Int64("product_id", request.ProductId), slog.Int64("status_id", int64(request.StatusId)))

	status, err := s.productSkStatusProvider.Get(ctx, request.ProductId, request.StatusId)
	if err != nil {
		if errors.Is(err, ErrProductStatusNotFound) {
			sl.Log.Warn("Product status not found", slog.Int64("product_id", request.ProductId), slog.Int64("status_id", int64(request.StatusId)), slog.String("op", op))
			return nil, err
		}
		sl.Log.Error("Failed to fetch product status", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Product status fetched successfully", slog.Int64("product_id", status.ProductId), slog.Int64("status_id", int64(request.StatusId)), slog.String("op", op))
	return status, nil
}
