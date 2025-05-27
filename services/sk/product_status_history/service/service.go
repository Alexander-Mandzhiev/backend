package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/product_status_history"
	"context"
	"errors"
	"log/slog"
)

var (
	ErrProductStatusNotFound = errors.New("product status not found")
	ErrInvalidRequest        = errors.New("invalid request")
)

type ProductSkStatusProvider interface {
	Create(ctx context.Context, req *product_status_history.CreateStatusRequest) (*product_status_history.ProductSkStatusResponse, error)
	GetByID(ctx context.Context, id int64) (*product_status_history.ProductSkStatusResponse, error)
	Delete(ctx context.Context, id int64) error
	ListByProduct(ctx context.Context, productID int64) ([]*product_status_history.ProductSkStatusResponse, error)
}

type Service struct {
	provider ProductSkStatusProvider
}

func New(provider ProductSkStatusProvider) *Service {
	op := "service.New"
	if provider == nil {
		sl.Log.Error("Status provider is nil", slog.String("op", op))
		return nil
	}

	sl.Log.Info("Service initialized", slog.String("op", op))
	return &Service{provider: provider}
}
