package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/locations"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, request *locations.DeleteLocationRequest) (*locations.DeleteLocationResponse, error) {
	op := "service.Delete"
	id := int(request.GetId())
	sl.Log.Debug("Deleting location", slog.String("op", op), slog.Int("id", id))

	err := s.locationProvider.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, ErrLocationNotFound) {
			sl.Log.Warn("Location not found during delete", slog.String("op", op), slog.Int("id", id))
		} else {
			sl.Log.Error("Failed to delete location", slog.String("op", op), slog.Any("error", err), slog.Int("id", id))
		}
		return &locations.DeleteLocationResponse{Success: false}, err
	}

	sl.Log.Info("Location deleted successfully", slog.String("op", op), slog.Int("id", id))
	return &locations.DeleteLocationResponse{Success: true}, nil
}
