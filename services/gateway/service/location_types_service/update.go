package location_types_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *location_types.UpdateLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	op := "LocationTypesService.Update"

	if req == nil || req.Id <= 0 {
		sl.Log.Warn("Invalid ID", slog.String("op", op))
		return nil, errors.New("invalid ID")
	}
	if req.Name == "" && req.Description == "" {
		sl.Log.Warn("No fields to update", slog.String("op", op))
		return nil, errors.New("no fields to update")
	}

	sl.Log.Info("Updating location type", slog.String("op", op), slog.Int("id", int(req.Id)), slog.String("name", req.Name), slog.String("description", req.Description))

	resp, err := s.client.Update(ctx, req)
	if err != nil {
		sl.Log.Error("Update failed", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	return resp, nil
}
