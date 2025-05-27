package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/product_status_history"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *product_status_history.CreateStatusRequest) (*product_status_history.ProductSkStatusResponse, error) {
	const op = "service.Create"
	logger := sl.Log.With(slog.String("op", op))

	if req == nil {
		logger.Error("nil request received")
		return nil, ErrInvalidRequest
	}

	logger.Debug("creating status", slog.Int64("product_id", req.ProductId), slog.Int("status_id", int(req.StatusId)))

	createdStatus, err := s.provider.Create(ctx, req)
	if err != nil {
		logger.Error("failed to create status", sl.Err(err), slog.Any("request", req))
		return nil, err
	}

	logger.Info("successfully created status", slog.Int64("record_id", createdStatus.Id),
		slog.Int64("product_id", createdStatus.ProductId), slog.Int("status_id", int(createdStatus.StatusId)))
	return createdStatus, nil
}
