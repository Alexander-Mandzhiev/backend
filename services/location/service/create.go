package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, request *locations.CreateLocationRequest) (*locations.LocationResponse, error) {
	op := "service.Create"
	newLocation := &locations.LocationResponse{
		Name:        request.GetName(),
		Type:        request.GetType(),
		Capacity:    request.GetCapacity(),
		CurrentLoad: request.GetCurrentLoad(),
	}

	sl.Log.Debug("Creating new location", slog.String("op", op), slog.Any("request", request))

	id, err := s.locationProvider.Create(ctx, newLocation)
	if err != nil {
		sl.Log.Error("Failed to create location", slog.String("op", op), slog.Any("error", err))
		return nil, err
	}

	newLocation.Id = id
	sl.Log.Info("Location created successfully", slog.String("op", op), slog.Int("id", int(id)))
	return newLocation, nil
}
