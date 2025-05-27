package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/products_sk"
	"context"
	"fmt"
	"log/slog"
	"math"
)

func (s *Service) Products(ctx context.Context, page, count int32) (*products_sk.ProductListResponse, error) {
	op := "service.List"
	sl.Log.Debug("Fetching all products", slog.String("op", op), slog.Int("page", int(page)), slog.Int("count", int(count)))

	if page < 1 || count < 1 {
		sl.Log.Warn("Invalid page or count value", slog.Int("page", int(page)), slog.Int("count", int(count)), slog.String("op", op))
		return nil, fmt.Errorf("page and count must be greater than zero")
	}

	products, totalItems, err := s.productSkProvider.Products(ctx, page, count)
	if err != nil {
		sl.Log.Error("Failed to fetch products", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	var totalPages int32
	if count > 0 {
		totalPages = int32(math.Ceil(float64(totalItems) / float64(count)))
	}

	sl.Log.Info("Products fetched successfully", slog.Int("count", len(products)), slog.String("op", op))
	return &products_sk.ProductListResponse{
		Data:       products,
		TotalPages: totalPages,
		TotalItems: totalItems,
	}, nil
}
