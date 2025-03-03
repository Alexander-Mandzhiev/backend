package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk_statuses"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, request *products_sk_statuses.ListProductStatusesRequest) (*products_sk_statuses.ProductStatusListResponse, error) {
	op := "service.List"
	sl.Log.Debug("Fetching all active product statuses", slog.String("op", op))

	statuses, err := s.productSkStatusProvider.List(ctx)
	if err != nil {
		sl.Log.Error("Failed to fetch product statuses", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Product statuses fetched successfully", slog.Int("count", len(statuses)), slog.String("op", op))
	return &products_sk_statuses.ProductStatusListResponse{Data: statuses}, nil
}
