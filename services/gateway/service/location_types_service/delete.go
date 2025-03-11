package location_types_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *location_types.DeleteLocationTypeRequest) (*location_types.DeleteLocationTypeResponse, error) {
	op := "LocationTypesService.Delete"
	if req.Id <= 0 {
		sl.Log.Warn("Invalid ID", slog.String("op", op))
		return nil, errors.New("invalid ID")
	}
	sl.Log.Info("Deleting location", slog.Int("id", int(req.Id)))

	resp, err := s.client.Delete(ctx, req)
	if err != nil {
		sl.Log.Error("Delete failed", sl.Err(err), slog.Int("id", int(req.Id)), slog.String("op", op))
		return nil, err
	}

	return resp, nil
}
