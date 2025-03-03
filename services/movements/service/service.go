package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/movements"
	"context"
	"errors"
	"log/slog"
)

var (
	ErrMovementNotFound = errors.New("movement not found")
)

type MovementsProvider interface {
	Create(ctx context.Context, movement *movements.MovementResponse) (int64, error)
	Movement(ctx context.Context, id int64) (*movements.MovementResponse, error)
	Update(ctx context.Context, movement *movements.MovementResponse) error
	Delete(ctx context.Context, id int64) error
	Movements(ctx context.Context, productId int64) ([]*movements.MovementResponse, error)
}

type Service struct {
	movementsProvider MovementsProvider
}

func New(movementsProvider MovementsProvider) *Service {
	op := "service.New"
	if movementsProvider == nil {
		sl.Log.Error("Location provider is nil", slog.String("op", op))
		return nil
	}

	sl.Log.Info("Service initialized", slog.String("op", op))
	return &Service{movementsProvider: movementsProvider}
}
