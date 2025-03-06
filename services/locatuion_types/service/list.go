package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, _ *location_types.ListLocationTypesRequest) (*location_types.LocationTypeListResponse, error) {
	op := "service.List"

	sl.Log.Debug("Fetching all location types", slog.String("op", op))

	locs, err := s.locationTypeProvider.List(ctx)
	if err != nil {
		sl.Log.Error("Failed to fetch location types", slog.String("op", op), slog.Any("error", err))
		return nil, err
	}

	sl.Log.Info("Location types fetched successfully", slog.String("op", op), slog.Int("count", len(locs)))
	return &location_types.LocationTypeListResponse{Data: locs}, nil
}
