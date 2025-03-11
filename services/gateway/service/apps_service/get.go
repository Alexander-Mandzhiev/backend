package apps_service

import (
	sl "backend/pkg/logger"
	app "backend/protos/gen/go/apps"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *app.GetAppRequest) (*app.GetAppResponse, error) {
	op := "apps.Get"
	if req.Id == 0 {
		sl.Log.Warn("Invalid ID", slog.String("op", op))
		return nil, errors.New("invalid ID")
	}

	sl.Log.Info("Fetching app", slog.Int("id", int(req.Id)), slog.String("op", op))

	resp, err := s.client.App(ctx, req)
	if err != nil {
		sl.Log.Error("Get failed", sl.Err(err), slog.Int("id", int(req.Id)), slog.String("op", op))
		return nil, err
	}

	return resp, nil
}
