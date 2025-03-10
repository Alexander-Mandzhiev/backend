package apps_service

import (
	sl "backend/pkg/logger"
	app "backend/protos/gen/go/apps"
	"context"
	"errors"
	"log/slog"
)

func (s *AppsService) Delete(ctx context.Context, req *app.DeleteRequest) (*app.DeleteResponse, error) {
	op := "apps.Delete"
	if req.Id == 0 {
		sl.Log.Warn("Invalid ID", slog.String("op", op))
		return nil, errors.New("invalid ID")
	}

	sl.Log.Info("Deleting app", slog.Int("id", int(req.Id)), slog.String("op", op))

	resp, err := s.client.Delete(ctx, req)
	if err != nil {
		sl.Log.Error("Delete failed", sl.Err(err), slog.Int("id", int(req.Id)), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("App deleted", slog.Int("id", int(req.Id)), slog.String("op", op))
	return resp, nil
}
