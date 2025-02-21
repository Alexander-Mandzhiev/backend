package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, request *locations.UpdateLocationRequest) (*locations.LocationResponse, error) {
	op := "service.Update"
	updatedLocation := &locations.LocationResponse{
		Id:          request.GetId(),
		Name:        request.GetName(),
		Type:        request.GetType(),
		Capacity:    request.GetCapacity(),
		CurrentLoad: request.GetCurrentLoad(),
	}

	sl.Log.Debug("Updating location", slog.String("op", op), slog.Any("request", request))

	err := s.locationProvider.Update(ctx, updatedLocation)
	if err != nil {
		if errors.Is(err, ErrLocationNotFound) {
			sl.Log.Warn("Location not found during update", slog.String("op", op), slog.Int("id", int(updatedLocation.Id)))
		} else {
			sl.Log.Error("Failed to update location", slog.String("op", op), slog.Any("error", err), slog.Int("id", int(updatedLocation.Id)))
		}
		return nil, err
	}

	sl.Log.Info("Location updated successfully", slog.String("op", op), slog.Int("id", int(updatedLocation.Id)))
	return updatedLocation, nil
}
