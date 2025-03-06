package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, request *location_types.GetLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	op := "service.Get"
	id := int(request.GetId())

	sl.Log.Debug("Fetching location type by ID", slog.String("op", op), slog.Int("id", id))

	loc, err := s.locationTypeProvider.Get(ctx, id)
	if err != nil {
		if errors.Is(err, ErrLocationTypesNotFound) {
			sl.Log.Warn("Location type not found", slog.String("op", op), slog.Int("id", id))
		} else {
			sl.Log.Error("Failed to fetch location type", slog.String("op", op), slog.Any("error", err), slog.Int("id", id))
		}
		return nil, err
	}

	sl.Log.Info("Location type fetched successfully", slog.String("op", op), slog.Int("id", int(loc.Id)))
	return loc, nil
}
