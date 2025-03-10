package statuses_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"context"
	"errors"
	"log/slog"
)

func (s *StatusesService) Update(ctx context.Context, req *statuses.UpdateStatusRequest) (*statuses.StatusResponse, error) {
	op := "statuses.Update"

	if req.Id == 0 {
		sl.Log.Warn("Invalid ID", slog.Int("id", int(req.Id)), slog.String("op", op))
		return nil, errors.New("invalid ID")
	}
	if req.Name == "" && req.Description == "" {
		sl.Log.Warn("No fields provided", slog.String("op", op))
		return nil, errors.New("at least one field must be provided")
	}

	sl.Log.Info("Updating status", slog.Int("id", int(req.Id)), slog.String("name", req.Name), slog.String("description", req.Description), slog.String("op", op))

	resp, err := s.client.Update(ctx, req)
	if err != nil {
		sl.Log.Error("Update failed", sl.Err(err), slog.Int("id", int(req.Id)), slog.String("op", op))
		return nil, err
	}

	return resp, nil
}
