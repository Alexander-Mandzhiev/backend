package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/movements"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, request *movements.UpdateMovementRequest) (*movements.MovementResponse, error) {
	op := "service.Update"
	sl.Log.Debug("Updating movement", slog.Int64("id", request.Id), slog.String("op", op))

	updatedMovement := &movements.MovementResponse{
		Id:             request.Id,
		ProductId:      request.ProductId,
		FromLocationId: request.FromLocationId,
		ToLocationId:   request.ToLocationId,
		UserId:         request.UserId,
		Comment:        request.Comment,
	}

	if err := s.movementsProvider.Update(ctx, updatedMovement); err != nil {
		sl.Log.Error("Failed to update movement", sl.Err(err), slog.Int64("id", request.Id), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Movement updated successfully", slog.Int64("id", request.Id), slog.String("op", op))
	return updatedMovement, nil
}
