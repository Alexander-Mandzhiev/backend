package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk_statuses"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, request *products_sk_statuses.CreateProductStatusRequest) (*products_sk_statuses.ProductSkStatusResponse, error) {
	op := "service.Create"
	sl.Log.Debug("Creating new product status", slog.String("op", op), slog.Int64("product_id", request.ProductId), slog.Int64("status_id", int64(request.StatusId)))

	newStatus := &products_sk_statuses.ProductSkStatusResponse{
		ProductId: request.ProductId,
		StatusId:  request.StatusId,
		Active:    true,
		CreatedAt: timestamppb.Now(),
	}

	err := s.productSkStatusProvider.Create(ctx, newStatus)
	if err != nil {
		sl.Log.Error("Failed to create product status", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Product status created successfully", slog.Int64("product_id", newStatus.ProductId), slog.Int64("status_id", int64(newStatus.StatusId)), slog.String("op", op))
	return newStatus, nil
}
