package location_types_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *location_types.GetLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	op := "LocationTypesService.Get"

	if req == nil || req.Id <= 0 {
		sl.Log.Warn("Invalid ID", slog.String("op", op))
		return nil, errors.New("invalid ID")
	}

	sl.Log.Info("Getting location type",
		slog.String("op", op),
		slog.Int("id", int(req.Id)),
	)

	resp, err := s.client.Get(ctx, req)
	if err != nil {
		sl.Log.Error("Get failed", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	return resp, nil
}
