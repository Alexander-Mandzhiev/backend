package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"backend/protos/gen/go/locations"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, request *locations.UpdateLocationRequest) (*locations.LocationResponse, error) {
	op := "service.Update"
	updatedLocation := &locations.UpdateLocationRequest{
		Id:          request.GetId(),
		Name:        request.GetName(),
		TypeId:      request.GetTypeId(),
		Capacity:    request.GetCapacity(),
		CurrentLoad: request.GetCurrentLoad(),
	}

	if updatedLocation.Capacity < 0 || updatedLocation.CurrentLoad < 0 {
		sl.Log.Warn("Invalid input data", slog.String("op", op), slog.Any("request", request))
		return nil, fmt.Errorf("negative values are not allowed for capacity or current_load")
	}

	locationType, err := s.locationTypesClient.Get(ctx, &location_types.GetLocationTypeRequest{Id: updatedLocation.TypeId})
	if err != nil {
		sl.Log.Warn("Location type not found", slog.String("op", op), slog.Int("type_id", int(updatedLocation.TypeId)))
		return nil, fmt.Errorf("location type with id %d not found", updatedLocation.TypeId)
	}
	sl.Log.Debug("Location type validated successfully", slog.String("op", op), slog.String("type_name", locationType.GetName()))

	if err = s.locationProvider.Update(ctx, updatedLocation); err != nil {
		if errors.Is(err, ErrLocationNotFound) {
			sl.Log.Warn("Location not found during update", slog.String("op", op), slog.Int("id", int(updatedLocation.Id)))
		} else {
			sl.Log.Error("Failed to update location", slog.String("op", op), slog.Any("error", err), slog.Int("id", int(updatedLocation.Id)))
		}
		return nil, err
	}

	sl.Log.Info("Location updated successfully", slog.String("op", op), slog.Int("id", int(updatedLocation.Id)))
	return &locations.LocationResponse{
		Id:          updatedLocation.Id,
		Name:        updatedLocation.Name,
		Type:        locationType.GetName(),
		Capacity:    updatedLocation.Capacity,
		CurrentLoad: updatedLocation.CurrentLoad,
	}, nil
}
