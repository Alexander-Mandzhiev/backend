package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Product(ctx context.Context, request *products_sk.GetProductRequest) (*products_sk.ProductResponse, error) {
	op := "service.Product"
	sl.Log.Debug("Fetching product by ID", slog.Int64("id", request.Id), slog.String("op", op))

	product, err := s.productSkProvider.Product(ctx, request.Id)
	if err != nil {
		if errors.Is(err, ErrProductSkNotFound) {
			sl.Log.Warn("Product not found", slog.Int64("id", request.Id), slog.String("op", op))
		} else {
			sl.Log.Error("Failed to fetch product", sl.Err(err), slog.Int64("id", request.Id), slog.String("op", op))
		}
		return nil, err
	}

	sl.Log.Info("Product fetched successfully", slog.Int64("id", product.Id), slog.String("op", op))
	return product, nil
}
