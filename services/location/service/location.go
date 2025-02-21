package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Location(ctx context.Context, request *locations.GetLocationRequest) (*locations.LocationResponse, error) {
	op := "service.Location"
	id := int(request.GetId())
	sl.Log.Debug("Fetching location by ID", slog.String("op", op), slog.Int("id", id))

	loc, err := s.locationProvider.Location(ctx, id)
	if err != nil {
		if errors.Is(err, ErrLocationNotFound) {
			sl.Log.Warn("Location not found", slog.String("op", op), slog.Int("id", id))
		} else {
			sl.Log.Error("Failed to fetch location", slog.String("op", op), slog.Any("error", err), slog.Int("id", id))
		}
		return nil, err
	}

	sl.Log.Info("Location fetched successfully", slog.String("op", op), slog.Int("id", int(loc.Id)))
	return loc, nil
}
