package location_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *locations.UpdateLocationRequest) (*locations.LocationResponse, error) {
	if req.Id <= 0 {
		sl.Log.Warn("Invalid ID", slog.String("op", "locations.Update"))
		return nil, errors.New("invalid ID")
	}

	if req.Name == "" && req.TypeId == 0 && req.Capacity == 0 {
		sl.Log.Warn("No fields provided", slog.String("op", "locations.Update"))
		return nil, errors.New("at least one field must be provided")
	}

	sl.Log.Info("Updating location", slog.Int("id", int(req.Id)), slog.String("name", req.Name), slog.Int("type_id", int(req.TypeId)), slog.Int("capacity", int(req.Capacity)))

	resp, err := s.client.Update(ctx, req)
	if err != nil {
		sl.Log.Error("Update failed", sl.Err(err), slog.Int("id", int(req.Id)), slog.String("op", "locations.Update"))
		return nil, err
	}

	return resp, nil
}
