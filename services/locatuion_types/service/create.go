package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, request *location_types.CreateLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	op := "service.Create"

	sl.Log.Debug("Creating new location type", slog.String("op", op), slog.Any("request", request))

	id, err := s.locationTypeProvider.Create(ctx, request)
	if err != nil {
		sl.Log.Error("Failed to create location type", slog.String("op", op), slog.Any("error", err))
		return nil, err
	}

	response := &location_types.LocationTypeResponse{
		Id:          id,
		Name:        request.GetName(),
		Description: request.GetDescription(),
	}

	sl.Log.Info("Location type created successfully", slog.String("op", op), slog.Int("id", int(id)))
	return response, nil
}
