package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/production_task"
	"context"
	"errors"
	"log/slog"
)

var (
	ErrInvalidCount = errors.New("count cannot be negative")
)

type ProductionTaskProvider interface {
	Tasks(ctx context.Context, params *production_task.RequestTaskParams) ([]*production_task.Product, error)
	TaskInPartName(ctx context.Context, params *production_task.RequestTaskParams) ([]*production_task.Product, error)
	RecordedInMsSQL(ctx context.Context, ids []int) error
	RecordedOutMsSQL(ctx context.Context, ids []int) error
}

type Service struct {
	productionTaskProvider ProductionTaskProvider
}

func New(productionTaskProvider ProductionTaskProvider) *Service {
	op := "service.New"
	if productionTaskProvider == nil {
		sl.Log.Error("Production Task provider is nil", slog.String("op", op))
		return nil
	}
	sl.Log.Info("Service initialized", slog.String("op", op))
	return &Service{productionTaskProvider: productionTaskProvider}
}
