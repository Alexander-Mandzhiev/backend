package location_types_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"log/slog"
)

func (s *LocationTypesService) List(ctx context.Context, req *location_types.ListLocationTypesRequest) (*location_types.LocationTypeListResponse, error) {
	op := "LocationTypesService.List"
	sl.Log.Info("Listing locations", slog.String("op", op))
	resp, err := s.client.List(ctx, req)
	if err != nil {
		sl.Log.Error("List location types failed", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	return resp, nil
}
