package statuses_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"context"
	"errors"
	"log/slog"
)

func (s *StatusesService) Get(ctx context.Context, req *statuses.GetStatusRequest) (*statuses.StatusResponse, error) {
	op := "statuses.Get"

	if req.Id == 0 {
		sl.Log.Warn("Invalid ID", slog.Int("id", int(req.Id)), slog.String("op", op))
		return nil, errors.New("invalid ID")
	}

	sl.Log.Info("Getting status", slog.Int("id", int(req.Id)), slog.String("op", op))
	resp, err := s.client.Status(ctx, req)
	if err != nil {
		sl.Log.Error("Get failed", sl.Err(err), slog.Int("id", int(req.Id)), slog.String("op", op))
		return nil, err
	}

	return resp, nil
}
