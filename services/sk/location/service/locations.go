package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/locations"

	"context"
	"log/slog"
)

func (s *Service) Locations(ctx context.Context) (*locations.LocationListResponse, error) {
	op := "service.Locations"
	sl.Log.Debug("Fetching all locations", slog.String("op", op))

	locsDB, err := s.locationProvider.Locations(ctx)
	if err != nil {
		sl.Log.Error("Failed to fetch locations", slog.String("op", op), slog.Any("error", err))
		return nil, err
	}

	var locs []*locations.LocationResponse
	for _, locDB := range locsDB {
		var typeName string
		if locDB.TypeId > 0 {
			locationType, err := s.locationTypesProvider.Get(ctx, int(locDB.TypeId))
			if err != nil {
				sl.Log.Warn("Location type not found", slog.String("op", op), slog.Int("type_id", int(locDB.TypeId)))
				typeName = ""
			} else {
				typeName = locationType.GetName()
			}
		}

		locs = append(locs, &locations.LocationResponse{
			Id:          locDB.Id,
			Name:        locDB.Name,
			Type:        typeName,
			Capacity:    locDB.Capacity,
			CurrentLoad: locDB.CurrentLoad,
		})
	}

	sl.Log.Info("Locations fetched successfully", slog.String("op", op), slog.Int("count", len(locs)))
	return &locations.LocationListResponse{Data: locs}, nil
}
