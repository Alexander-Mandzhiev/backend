package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/product_status_history"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, request *product_status_history.DeleteStatusRequest) (*product_status_history.DeleteStatusResponse, error) {
	op := "service.Delete"
	sl.Log.Debug("Deactivating product status", slog.String("op", op), slog.String("op", op), slog.Int64("id", request.GetId()))
	err := s.provider.Delete(ctx, request.GetId())
	if err != nil {
		if errors.Is(err, ErrProductStatusNotFound) {
			sl.Log.Warn("Product status not found for deactivation", slog.Int64("id", request.GetId()), slog.String("op", op))
			return nil, err
		}
		sl.Log.Error("Failed to deactivate product status", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Product status deactivated successfully", slog.Int64("id", request.GetId()), slog.String("op", op))
	return &product_status_history.DeleteStatusResponse{Success: true}, nil
}
