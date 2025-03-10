package location_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"errors"
	"log/slog"
)

func (s *LocationService) Create(ctx context.Context, req *locations.CreateLocationRequest) (*locations.LocationResponse, error) {
	if req.Name == "" || req.TypeId <= 0 || req.Capacity <= 0 {
		sl.Log.Warn("Missing required fields", slog.String("op", "locations.Create"))
		return nil, errors.New("name, type_id, and capacity are required")
	}

	sl.Log.Info("Creating location", slog.String("name", req.Name), slog.Int("type_id", int(req.TypeId)), slog.Int("capacity", int(req.Capacity)))

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		sl.Log.Error("Create failed", sl.Err(err), slog.String("op", "locations.Create"))
		return nil, err
	}

	return resp, nil
}
