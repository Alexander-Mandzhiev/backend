package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/locations"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, request *locations.CreateLocationRequest) (*locations.LocationResponse, error) {
	op := "service.Create"
	newLocation := &locations.CreateLocationRequest{
		Name:        request.GetName(),
		TypeId:      request.GetTypeId(),
		Capacity:    request.GetCapacity(),
		CurrentLoad: request.GetCurrentLoad(),
	}

	if newLocation.Capacity < 0 || newLocation.CurrentLoad < 0 {
		sl.Log.Warn("Invalid input data", slog.String("op", op), slog.Any("request", request))
		return nil, fmt.Errorf("negative values are not allowed for capacity or current_load")
	}

	locationType, err := s.locationTypesProvider.Get(ctx, int(request.GetTypeId()))
	if err != nil {
		sl.Log.Warn("Location type not found", slog.String("op", op), slog.Int("type_id", int(request.GetTypeId())))
		return nil, fmt.Errorf("location type with id %d not found", request.GetTypeId())
	}

	id, err := s.locationProvider.Create(ctx, newLocation)
	if err != nil {
		sl.Log.Error("Failed to create location", slog.String("op", op), slog.Any("error", err))
		return nil, err
	}

	sl.Log.Info("Location created successfully", slog.String("op", op), slog.Int("id", int(id)))
	return &locations.LocationResponse{
		Id:          id,
		Name:        request.GetName(),
		Type:        locationType.Name,
		Capacity:    request.GetCapacity(),
		CurrentLoad: request.GetCurrentLoad(),
	}, nil
}
