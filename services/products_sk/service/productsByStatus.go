package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk"
	"context"
	"log/slog"
)

func (s *Service) ProductsByStatus(ctx context.Context, statusID int32) (*products_sk.ProductListResponse, error) {
	op := "service.ProductsByStatus"
	sl.Log.Debug("Fetching products by active status", slog.Int64("status_id", int64(statusID)), slog.String("op", op))

	products, err := s.productSkProvider.ProductsByStatus(ctx, statusID)
	if err != nil {
		sl.Log.Error("Failed to fetch products by status", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Products fetched successfully", slog.Int("count", len(products)), slog.String("op", op))
	return &products_sk.ProductListResponse{Data: products}, nil
}
