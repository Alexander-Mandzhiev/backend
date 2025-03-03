package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk_statuses"
	"context"
	"errors"
	"log/slog"
)

var ErrProductStatusNotFound = errors.New("product status not found")

type ProductSkStatusProvider interface {
	Create(ctx context.Context, productStatus *products_sk_statuses.ProductSkStatusResponse) error
	Get(ctx context.Context, productID int64, statusID int32) (*products_sk_statuses.ProductSkStatusResponse, error)
	Update(ctx context.Context, productStatus *products_sk_statuses.ProductSkStatusResponse) error
	Delete(ctx context.Context, productID int64, statusID int32) error
	List(ctx context.Context) ([]*products_sk_statuses.ProductSkStatusResponse, error)
}

type Service struct {
	productSkStatusProvider ProductSkStatusProvider
}

func New(productSkStatusProvider ProductSkStatusProvider) *Service {
	op := "service.New"
	if productSkStatusProvider == nil {
		sl.Log.Error("Location provider is nil", slog.String("op", op))
		return nil
	}

	sl.Log.Info("Service initialized", slog.String("op", op))
	return &Service{productSkStatusProvider: productSkStatusProvider}
}
