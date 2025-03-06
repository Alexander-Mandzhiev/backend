package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, request *location_types.UpdateLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	op := "service.Update"

	sl.Log.Debug("Updating location type", slog.String("op", op), slog.Any("request", request))

	err := s.locationTypeProvider.Update(ctx, request)
	if err != nil {
		if errors.Is(err, ErrLocationTypesNotFound) {
			sl.Log.Warn("Location type not found during update", slog.String("op", op), slog.Int("id", int(request.GetId())))
		} else {
			sl.Log.Error("Failed to update location type", slog.String("op", op), slog.Any("error", err), slog.Int("id", int(request.GetId())))
		}
		return nil, err
	}

	response := &location_types.LocationTypeResponse{
		Id:          request.GetId(),
		Name:        request.GetName(),
		Description: request.GetDescription(),
	}

	sl.Log.Info("Location type updated successfully", slog.String("op", op), slog.Int("id", int(request.GetId())))
	return response, nil
}
