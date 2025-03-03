package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk_statuses"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, request *products_sk_statuses.DeleteProductStatusRequest) (*products_sk_statuses.DeleteProductStatusResponse, error) {
	op := "service.Delete"
	sl.Log.Debug("Deactivating product status", slog.String("op", op), slog.Int64("product_id", request.ProductId), slog.Int64("status_id", int64(request.StatusId)))
	err := s.productSkStatusProvider.Delete(ctx, request.ProductId, request.StatusId)
	if err != nil {
		if errors.Is(err, ErrProductStatusNotFound) {
			sl.Log.Warn("Product status not found for deactivation", slog.Int64("product_id", request.ProductId), slog.Int64("status_id", int64(request.StatusId)), slog.String("op", op))
			return nil, err
		}
		sl.Log.Error("Failed to deactivate product status", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Product status deactivated successfully", slog.Int64("product_id", request.ProductId), slog.Int64("status_id", int64(request.StatusId)), slog.String("op", op))
	return &products_sk_statuses.DeleteProductStatusResponse{Success: true}, nil
}
