package location_types_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *location_types.CreateLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	op := "LocationTypesService.Create"

	if req == nil {
		sl.Log.Error("Empty request", slog.String("op", op))
		return nil, errors.New("empty request")
	}
	if req.Name == "" {
		sl.Log.Warn("Missing name field", slog.String("op", op))
		return nil, errors.New("name is required")
	}

	sl.Log.Info("Creating location type", slog.String("op", op), slog.String("name", req.Name), slog.String("description", req.Description))

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		sl.Log.Error("Create failed", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Location type created", slog.String("op", op), slog.Int("id", int(resp.Id)))

	return resp, nil
}
