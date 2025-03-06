package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"backend/protos/gen/go/locations"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Location(ctx context.Context, request *locations.GetLocationRequest) (*locations.LocationResponse, error) {
	op := "service.Location"
	id := int(request.GetId())
	sl.Log.Debug("Fetching location by ID", slog.String("op", op), slog.Int("id", id))

	locDB, err := s.locationProvider.Location(ctx, id)
	if err != nil {
		if errors.Is(err, ErrLocationNotFound) {
			sl.Log.Warn("Location not found", slog.String("op", op), slog.Int("id", id))
		} else {
			sl.Log.Error("Failed to fetch location", slog.String("op", op), slog.Any("error", err), slog.Int("id", id))
		}
		return nil, err
	}

	var typeName string
	if locDB.TypeId > 0 {
		locationType, err := s.locationTypesClient.Get(ctx, &location_types.GetLocationTypeRequest{Id: locDB.TypeId})
		if err != nil {
			sl.Log.Warn("Location type not found", slog.String("op", op), slog.Int("type_id", int(locDB.TypeId)))
			typeName = ""
		} else {
			typeName = locationType.GetName()
		}
	}

	sl.Log.Info("Location fetched successfully", slog.String("op", op), slog.Int("id", int(locDB.Id)))
	return &locations.LocationResponse{
		Id:          locDB.Id,
		Name:        locDB.Name,
		Type:        typeName,
		Capacity:    locDB.Capacity,
		CurrentLoad: locDB.CurrentLoad,
	}, nil
}
