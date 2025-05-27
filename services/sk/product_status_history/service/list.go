package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/product_status_history"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *product_status_history.ListStatusesByProductRequest) (*product_status_history.StatusListResponse, error) {
	op := "service.List"
	logger := sl.Log.With(slog.String("op", op))

	if req == nil {
		logger.Error("nil request received")
		return nil, ErrInvalidRequest
	}

	statuses, err := s.provider.ListByProduct(ctx, req.ProductId)
	if err != nil {
		logger.Error("failed to list statuses", sl.Err(err), slog.Int64("product_id", req.ProductId))
		return nil, err
	}

	return &product_status_history.StatusListResponse{Statuses: statuses}, nil
}
