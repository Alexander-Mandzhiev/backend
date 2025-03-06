package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, request *location_types.DeleteLocationTypeRequest) (*location_types.DeleteLocationTypeResponse, error) {
	op := "service.Delete"
	id := int(request.GetId())

	sl.Log.Debug("Deleting location type", slog.String("op", op), slog.Int("id", id))

	err := s.locationTypeProvider.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, ErrLocationTypesNotFound) {
			sl.Log.Warn("Location type not found during delete", slog.String("op", op), slog.Int("id", id))
		} else {
			sl.Log.Error("Failed to delete location type", slog.String("op", op), slog.Any("error", err), slog.Int("id", id))
		}
		return &location_types.DeleteLocationTypeResponse{Success: false}, err
	}

	sl.Log.Info("Location type deleted successfully", slog.String("op", op), slog.Int("id", id))
	return &location_types.DeleteLocationTypeResponse{Success: true}, nil
}
