package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"context"
	"errors"
	"log/slog"
)

var (
	ErrStatusNotFound = errors.New("status not found")
)

type StatusesProvider interface {
	Create(ctx context.Context, status *statuses.StatusResponse) (int, error)
	Status(ctx context.Context, id int) (*statuses.StatusResponse, error)
	Update(ctx context.Context, status *statuses.StatusResponse) error
	Delete(ctx context.Context, id int) error
	Statuses(ctx context.Context) ([]*statuses.StatusResponse, error)
}

type Service struct {
	statusesProvider StatusesProvider
}

func New(statusesProvider StatusesProvider) *Service {
	op := "service.New"
	if statusesProvider == nil {
		sl.Log.Error("Location provider is nil", slog.String("op", op))
		return nil
	}

	sl.Log.Info("Service initialized", slog.String("op", op))
	return &Service{statusesProvider: statusesProvider}
}
