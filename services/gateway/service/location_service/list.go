package location_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"log/slog"
)

func (s *LocationService) List(ctx context.Context, req *locations.ListLocationsRequest) (*locations.LocationListResponse, error) {
	op := "LocationService.List"
	sl.Log.Info("Listing locations", slog.String("op", op))
	resp, err := s.client.List(ctx, req)
	if err != nil {
		sl.Log.Error("List failed", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	return resp, nil
}
