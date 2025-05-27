package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/statuses"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, request *statuses.CreateStatusRequest) (*statuses.StatusResponse, error) {
	const op = "service.Create"
	const maxNameLength = 255

	if request.GetName() == "" {
		sl.Log.Warn("Empty status name", slog.String("op", op), slog.Any("request", request))
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	if len(request.GetName()) > maxNameLength {
		sl.Log.Warn("Status name too long", slog.String("op", op), slog.Int("max_length", maxNameLength), slog.Int("received_length", len(request.GetName())))
		return nil, status.Error(codes.InvalidArgument, "name must be less than 255 characters")
	}

	sl.Log.Debug("Creating new status", slog.String("op", op), slog.String("name", request.GetName()))

	id, err := s.statusesProvider.Create(ctx, &statuses.StatusResponse{
		Name:        request.GetName(),
		Description: request.GetDescription(),
	})
	if err != nil {
		sl.Log.Error("Create operation failed", slog.String("op", op), slog.Any("error", err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	response := &statuses.StatusResponse{
		Id:          int32(id),
		Name:        request.GetName(),
		Description: request.GetDescription(),
	}

	sl.Log.Info("Status created", slog.String("op", op), slog.Int("id", id), slog.String("name", response.GetName()))
	return response, nil
}
