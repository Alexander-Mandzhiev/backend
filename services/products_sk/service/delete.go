package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, request *products_sk.DeleteProductRequest) (*products_sk.DeleteProductResponse, error) {
	op := "service.Delete"
	sl.Log.Debug("Deleting product", slog.Int64("id", request.Id), slog.String("op", op))

	if err := s.productSkProvider.Delete(ctx, request.Id); err != nil {
		if errors.Is(err, ErrProductSkNotFound) {
			sl.Log.Warn("Product not found for deletion", slog.Int64("id", request.Id), slog.String("op", op))
		} else {
			sl.Log.Error("Failed to delete product", sl.Err(err), slog.Int64("id", request.Id), slog.String("op", op))
		}
		return nil, err
	}

	sl.Log.Info("Product deleted successfully", slog.Int64("id", request.Id), slog.String("op", op))
	return &products_sk.DeleteProductResponse{Success: true}, nil
}
