package statuses_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"context"
	"errors"
	"log/slog"
)

func (s *StatusesService) Delete(ctx context.Context, req *statuses.DeleteStatusRequest) (*statuses.DeleteStatusResponse, error) {
	op := "statuses.Delete"
	if req.Id == 0 {
		sl.Log.Warn("Invalid ID", slog.Int("id", int(req.Id)), slog.String("op", op))
		return nil, errors.New("invalid ID")
	}

	sl.Log.Info("Deleting status", slog.Int("id", int(req.Id)), slog.String("op", op))
	resp, err := s.client.Delete(ctx, req)
	if err != nil {
		sl.Log.Error("Delete failed", sl.Err(err), slog.Int("id", int(req.Id)), slog.String("op", op))
		return nil, err
	}

	return resp, nil
}
