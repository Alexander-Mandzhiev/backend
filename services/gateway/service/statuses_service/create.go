package statuses_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"context"
	"errors"
	"log/slog"
)

func (s *StatusesService) Create(ctx context.Context, req *statuses.CreateStatusRequest) (*statuses.StatusResponse, error) {
	op := "statuses.Create"

	if req.Name == "" {
		sl.Log.Warn("Name is required", slog.String("op", op))
		return nil, errors.New("name is required")
	}

	sl.Log.Info("Creating status", slog.String("name", req.Name), slog.String("description", req.Description), slog.String("op", op))

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		sl.Log.Error("Create failed", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Status created", slog.Int("id", int(resp.Id)), slog.String("op", op))
	return resp, nil
}
