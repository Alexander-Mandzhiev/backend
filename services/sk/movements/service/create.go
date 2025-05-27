package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/movements"
	"backend/protos/gen/go/sk/product_status_history"
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, request *movements.CreateMovementRequest) (*movements.MovementResponse, error) {
	const op = "movements.Service.Create"
	logger := sl.Log.With(slog.String("op", op))

	if request == nil {
		logger.Error("nil request")
		return nil, errors.New("request cannot be nil")
	}

	product, err := s.productSkProvider.Product(ctx, request.ProductId)
	if err != nil {
		logger.Error("product not found",
			slog.Int64("product_id", request.ProductId),
			sl.Err(err))
		return nil, fmt.Errorf("product %d not found: %w", request.ProductId, err)
	}

	newMovement := &movements.MovementResponse{
		ProductId:      product.Id,
		FromLocationId: request.FromLocationId,
		ToLocationId:   request.ToLocationId,
		UserId:         request.UserId,
		Comment:        request.Comment,
		CreatedAt:      timestamppb.Now(),
	}

	id, err := s.movementsProvider.Create(ctx, newMovement)
	if err != nil {
		logger.Error("failed to create movement", slog.Any("request", request), sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	newMovement.Id = id

	statusReq := &product_status_history.CreateStatusRequest{
		ProductId: request.ProductId,
		StatusId:  request.GetStatusId(),
	}

	if _, err = s.productStatusHistory.Create(ctx, statusReq); err != nil {
		logger.Error("failed to update status", slog.Int64("product_id", request.ProductId), sl.Err(err))

		if delErr := s.movementsProvider.Delete(ctx, id); delErr != nil {
			logger.Error("rollback failed", slog.Int64("movement_id", id), sl.Err(delErr))
		}

		return nil, fmt.Errorf("status update failed: %w", err)
	}

	logger.Info("movement created", slog.Int64("id", id), slog.Int64("product_id", request.ProductId))
	return newMovement, nil
}
