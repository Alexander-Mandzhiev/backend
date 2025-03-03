package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk"
	"context"
	"log/slog"
)

func (s *Service) Products(ctx context.Context) (*products_sk.ProductListResponse, error) {
	op := "service.List"
	sl.Log.Debug("Fetching all products", slog.String("op", op))

	products, err := s.productSkProvider.Products(ctx)
	if err != nil {
		sl.Log.Error("Failed to fetch products", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Products fetched successfully", slog.Int("count", len(products)), slog.String("op", op))
	return &products_sk.ProductListResponse{Data: products}, nil
}
