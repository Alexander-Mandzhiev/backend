package apps_service

import (
	sl "backend/pkg/logger"
	app "backend/protos/gen/go/apps"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *app.UpdateRequest) (*app.UpdateResponse, error) {
	op := "apps.Update"
	if req.Id == 0 {
		sl.Log.Warn("Invalid ID", slog.String("op", op))
		return nil, errors.New("invalid ID")
	}
	if req.Name == "" {
		sl.Log.Warn("Name is required", slog.String("op", op))
		return nil, errors.New("name is required")
	}

	sl.Log.Info("Updating app", slog.Int("id", int(req.Id)), slog.String("name", req.Name), slog.String("op", op))

	resp, err := s.client.Update(ctx, req)
	if err != nil {
		sl.Log.Error("Update failed", sl.Err(err), slog.Int("id", int(req.Id)), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("App updated", slog.Int("id", int(req.Id)), slog.String("op", op))
	return resp, nil
}
