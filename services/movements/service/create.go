package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/movements"
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, request *movements.CreateMovementRequest) (*movements.MovementResponse, error) {
	op := "service.Create"
	sl.Log.Debug("Creating new movement", slog.String("op", op))
	if request == nil {
		return nil, errors.New("request cannot be nil")
	}

	newMovement := &movements.MovementResponse{
		ProductId:      request.ProductId,
		FromLocationId: request.FromLocationId,
		ToLocationId:   request.ToLocationId,
		UserId:         request.UserId,
		Comment:        request.Comment,
		CreatedAt:      timestamppb.Now(),
	}

	id, err := s.movementsProvider.Create(ctx, newMovement)
	if err != nil {
		sl.Log.Error("Failed to create movement", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	newMovement.Id = id
	sl.Log.Info("Movement created successfully", slog.Int64("id", id), slog.String("op", op))
	return newMovement, nil
}
