package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/location_types"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, request *location_types.GetLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	const op = "service.LocationTypes.Get"
	id := int(request.GetId())

	if id <= 0 {
		sl.Log.Warn("Invalid ID format", slog.String("op", op), slog.Int("id", id))
		return nil, status.Error(codes.InvalidArgument, "ID must be a positive integer")
	}

	sl.Log.Debug("Fetching location type", slog.String("op", op), slog.Int("id", id))

	response, err := s.locationTypeProvider.Get(ctx, id)
	if err != nil {
		if errors.Is(err, ErrLocationTypesNotFound) {
			sl.Log.Warn("Location type not found", slog.String("op", op), slog.Int("id", id))
			return nil, status.Error(codes.NotFound, "location type not found")
		}

		sl.Log.Error("Get failed", slog.String("op", op), slog.Int("id", id), slog.Any("error", err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	sl.Log.Info("Location type fetched", slog.String("op", op), slog.Int("id", id), slog.String("name", response.GetName()))
	return response, nil
}
