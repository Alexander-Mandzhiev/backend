package location_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *locations.DeleteLocationRequest) (*locations.DeleteLocationResponse, error) {
	if req.Id <= 0 {
		sl.Log.Warn("Invalid ID", slog.String("op", "locations.Delete"))
		return nil, errors.New("invalid ID")
	}

	sl.Log.Info("Deleting location", slog.Int("id", int(req.Id)))
	resp, err := s.client.Delete(ctx, req)
	if err != nil {
		sl.Log.Error("Delete failed", sl.Err(err), slog.Int("id", int(req.Id)), slog.String("op", "locations.Delete"))
		return nil, err
	}

	return resp, nil
}
