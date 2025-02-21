package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"errors"
	"log/slog"
)

var (
	ErrLocationNotFound = errors.New("location not found")
)

type LocationProvider interface {
	Create(ctx context.Context, location *locations.LocationResponse) (int32, error)
	Locations(ctx context.Context) ([]*locations.LocationResponse, error)
	Location(ctx context.Context, id int) (*locations.LocationResponse, error)
	Update(ctx context.Context, location *locations.LocationResponse) error
	Delete(ctx context.Context, id int) error
}

type Service struct {
	locationProvider LocationProvider
}

func New(locationProvider LocationProvider) *Service {
	op := "service.New"
	if locationProvider == nil {
		sl.Log.Error("Location provider is nil", slog.String("op", op))
		return nil
	}

	sl.Log.Info("Service initialized", slog.String("op", op))
	return &Service{locationProvider: locationProvider}
}
