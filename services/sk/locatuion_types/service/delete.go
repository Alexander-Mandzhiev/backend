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

func (s *Service) Delete(ctx context.Context, request *location_types.DeleteLocationTypeRequest) (*location_types.DeleteLocationTypeResponse, error) {
	const op = "service.LocationTypes.Delete"
	id := int(request.GetId())

	if id <= 0 {
		sl.Log.Warn("Invalid ID format", slog.String("op", op), slog.Int("id", id))
		return nil, status.Error(codes.InvalidArgument, "ID must be a positive integer")
	}

	sl.Log.Debug("Deleting location type", slog.String("op", op), slog.Int("id", id))

	err := s.locationTypeProvider.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, ErrLocationTypesNotFound) {
			sl.Log.Warn("Location type not found", slog.String("op", op), slog.Int("id", id))
			return &location_types.DeleteLocationTypeResponse{Success: false}, status.Error(codes.NotFound, "location type not found")
		}

		sl.Log.Error("Delete failed", slog.String("op", op), slog.Int("id", id), slog.Any("error", err))
		return &location_types.DeleteLocationTypeResponse{Success: false}, status.Error(codes.Internal, "internal server error")
	}

	sl.Log.Info("Location type deleted", slog.String("op", op), slog.Int("id", id))
	return &location_types.DeleteLocationTypeResponse{Success: true}, nil
}
