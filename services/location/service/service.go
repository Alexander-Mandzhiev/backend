package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"backend/protos/gen/go/locations"
	"context"
	"errors"
	"log/slog"
)

var (
	ErrLocationNotFound = errors.New("location not found")
)

type LocationProvider interface {
	Create(ctx context.Context, request *locations.CreateLocationRequest) (int32, error)
	Locations(ctx context.Context) ([]*locations.UpdateLocationRequest, error)
	Location(ctx context.Context, id int) (*locations.UpdateLocationRequest, error)
	Update(ctx context.Context, request *locations.UpdateLocationRequest) error
	Delete(ctx context.Context, id int) error
}

type Service struct {
	locationProvider    LocationProvider
	locationTypesClient location_types.LocationTypeServiceClient
}

func New(locationProvider LocationProvider, locationTypesClient location_types.LocationTypeServiceClient) *Service {
	op := "service.New"
	if locationProvider == nil {
		sl.Log.Error("Location provider is nil", slog.String("op", op))
		return nil
	}

	sl.Log.Info("Service initialized", slog.String("op", op))
	return &Service{locationProvider: locationProvider, locationTypesClient: locationTypesClient}
}
