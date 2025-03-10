package location_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"errors"
	"log/slog"
)

func (s *LocationService) Get(ctx context.Context, req *locations.GetLocationRequest) (*locations.LocationResponse, error) {
	if req.Id <= 0 {
		sl.Log.Warn("Invalid ID", slog.String("op", "locations.Get"))
		return nil, errors.New("invalid ID")
	}

	sl.Log.Info("Getting location", slog.Int("id", int(req.Id)))
	resp, err := s.client.Location(ctx, req)
	if err != nil {
		sl.Log.Error("Get failed", sl.Err(err), slog.Int("id", int(req.Id)), slog.String("op", "locations.Get"))
		return nil, err
	}

	return resp, nil
}
