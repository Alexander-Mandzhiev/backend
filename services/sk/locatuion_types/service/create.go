package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/location_types"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, request *location_types.CreateLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	const op = "service.LocationTypes.Create"
	const maxNameLength = 255

	if request.GetName() == "" {
		sl.Log.Warn("Empty location type name", slog.String("op", op))
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	if len(request.GetName()) > maxNameLength {
		sl.Log.Warn("Name exceeds maximum length", slog.String("op", op), slog.Int("max_length", maxNameLength),
			slog.Int("current_length", len(request.GetName())))
		return nil, status.Error(codes.InvalidArgument, "name must be less than 255 characters")
	}

	sl.Log.Debug("Creating new location type", slog.String("op", op), slog.Any("request", request))

	response, err := s.locationTypeProvider.Create(ctx, request)
	if err != nil {
		sl.Log.Error("Failed to create location type", slog.String("op", op), slog.Any("error", err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	sl.Log.Info("Location type created successfully", slog.String("op", op), slog.Int("id", int(response.Id)))
	return response, nil
}
