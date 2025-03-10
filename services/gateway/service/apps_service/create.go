package apps_service

import (
	sl "backend/pkg/logger"
	app "backend/protos/gen/go/apps"
	"context"
	"errors"
	"log/slog"
)

func (s *AppsService) Create(ctx context.Context, req *app.CreateRequest) (*app.CreateResponse, error) {
	op := "apps.Create"
	if req.Name == "" {
		sl.Log.Warn("Name is required", slog.String("op", op))
		return nil, errors.New("name is required")
	}

	sl.Log.Info("Creating app", slog.String("name", req.Name), slog.String("op", op))

	resp, err := s.client.Create(ctx, req)
	if err != nil {
		sl.Log.Error("Create failed", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("App created", slog.Int("id", int(resp.Id)), slog.String("op", op))
	return resp, nil
}
