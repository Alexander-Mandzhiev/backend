package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/location_types"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) List(ctx context.Context, _ *location_types.ListLocationTypesRequest) (*location_types.LocationTypeListResponse, error) {
	const op = "service.LocationTypes.List"

	sl.Log.Debug("Fetching all location types", slog.String("op", op))

	items, err := s.locationTypeProvider.List(ctx)
	if err != nil {
		sl.Log.Error("Failed to fetch location types", slog.String("op", op), slog.Any("error", err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	response := &location_types.LocationTypeListResponse{Data: items}

	sl.Log.Info("Location types fetched successfully", slog.String("op", op), slog.Int("count", len(items)))
	return response, nil
}
