package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk_statuses"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, request *products_sk_statuses.UpdateProductStatusRequest) (*products_sk_statuses.ProductSkStatusResponse, error) {
	op := "service.Update"
	sl.Log.Debug("Updating product status", slog.String("op", op), slog.Int64("product_id", request.ProductId), slog.Int64("status_id", int64(request.StatusId)))

	existingStatus, err := s.productSkStatusProvider.Get(ctx, request.ProductId, request.StatusId)
	if err != nil {
		if errors.Is(err, ErrProductStatusNotFound) {
			sl.Log.Warn("Product status not found for update", slog.Int64("product_id", request.ProductId), slog.Int64("status_id", int64(request.StatusId)), slog.String("op", op))
			return nil, err
		}
		sl.Log.Error("Failed to fetch product status for update", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	existingStatus.Active = request.Active
	err = s.productSkStatusProvider.Update(ctx, existingStatus)
	if err != nil {
		sl.Log.Error("Failed to update product status", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Product status updated successfully", slog.Int64("product_id", existingStatus.ProductId), slog.Int64("status_id", int64(existingStatus.StatusId)), slog.String("op", op))
	return existingStatus, nil
}
