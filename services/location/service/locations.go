package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"log/slog"
)

func (s *Service) Locations(ctx context.Context) (*locations.LocationListResponse, error) {
	op := "service.Locations"
	sl.Log.Debug("Fetching all locations", slog.String("op", op))

	locs, err := s.locationProvider.Locations(ctx)
	if err != nil {
		sl.Log.Error("Failed to fetch locations", slog.String("op", op), slog.Any("error", err))
		return nil, err
	}

	sl.Log.Info("Locations fetched successfully", slog.String("op", op), slog.Int("count", len(locs)))

	return &locations.LocationListResponse{Data: locs}, nil
}
