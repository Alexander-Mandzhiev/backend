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

func (s *Service) Update(ctx context.Context, request *location_types.UpdateLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	const op = "service.LocationTypes.Update"
	const maxNameLength = 255

	if request.GetId() <= 0 {
		sl.Log.Warn("Invalid ID format", slog.String("op", op), slog.Int("id", int(request.GetId())))
		return nil, status.Error(codes.InvalidArgument, "ID must be a positive integer")
	}

	if len(request.GetName()) > maxNameLength {
		sl.Log.Warn("Name exceeds maximum length", slog.String("op", op), slog.Int("max_length", maxNameLength), slog.Int("current_length", len(request.GetName())))
		return nil, status.Error(codes.InvalidArgument, "name too long")
	}

	sl.Log.Debug("Updating location type", slog.String("op", op), slog.Int("id", int(request.GetId())), slog.String("new_name", request.GetName()))

	updatedLoc, err := s.locationTypeProvider.Update(ctx, request)
	if err != nil {
		if errors.Is(err, ErrLocationTypesNotFound) {
			sl.Log.Warn("Location type not found", slog.String("op", op), slog.Int("id", int(request.GetId())))
			return nil, status.Error(codes.NotFound, "location type not found")
		}

		sl.Log.Error("Update failed", slog.String("op", op), slog.Int("id", int(request.GetId())), slog.Any("error", err))
		return nil, status.Error(codes.Internal, "update failed")
	}

	sl.Log.Info("Location type updated", slog.String("op", op), slog.Int("id", int(updatedLoc.GetId())), slog.String("new_name", updatedLoc.GetName()))

	return updatedLoc, nil
}
