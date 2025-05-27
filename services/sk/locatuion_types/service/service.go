package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/location_types"
	"context"
	"errors"
	"log/slog"
)

var (
	ErrLocationTypesNotFound = errors.New("location_types not found")
)

type LocationTypeProvider interface {
	Create(ctx context.Context, locationTypes *location_types.CreateLocationTypeRequest) (*location_types.LocationTypeResponse, error)
	Get(ctx context.Context, id int) (*location_types.LocationTypeResponse, error)
	Update(ctx context.Context, locationTypes *location_types.UpdateLocationTypeRequest) (*location_types.LocationTypeResponse, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*location_types.LocationTypeResponse, error)
}

type Service struct {
	locationTypeProvider LocationTypeProvider
}

func New(locationTypeProvider LocationTypeProvider) *Service {
	op := "service.New"
	if locationTypeProvider == nil {
		sl.Log.Error("Location provider is nil", slog.String("op", op))
		return nil
	}

	sl.Log.Info("Service initialized", slog.String("op", op))
	return &Service{locationTypeProvider: locationTypeProvider}
}
