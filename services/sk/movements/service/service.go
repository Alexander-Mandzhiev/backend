package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/movements"
	serviceL "backend/services/sk/location/service"
	servicePSH "backend/services/sk/product_status_history/service"
	serviceS "backend/services/sk/statuses/service"
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
	InitializeProduct(ctx context.Context, req *movements.InitializeProductRequest) (*movements.InitializeProductResponse, error)
}

type Service struct {
	movementsProvider    MovementsProvider
	statusesProvider     serviceS.StatusesProvider
	locationProvider     serviceL.LocationProvider
	productStatusHistory servicePSH.ProductSkStatusProvider
}

func New(movementsProvider MovementsProvider, statusesProvider serviceS.StatusesProvider,
	productStatusHistory servicePSH.ProductSkStatusProvider, locationProvider serviceL.LocationProvider) *Service {
	op := "service.New"
	if movementsProvider == nil {
		sl.Log.Error("Location provider is nil", slog.String("op", op))
		return nil
	}

	sl.Log.Info("Service initialized", slog.String("op", op))
	return &Service{movementsProvider: movementsProvider, statusesProvider: statusesProvider,
		productStatusHistory: productStatusHistory, locationProvider: locationProvider}
}
