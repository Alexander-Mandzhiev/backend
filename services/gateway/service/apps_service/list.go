package apps_service

import (
	sl "backend/pkg/logger"
	app "backend/protos/gen/go/apps"
	"context"
	"log/slog"
)

func (s *AppsService) List(ctx context.Context, req *app.GetAppsRequest) (*app.GetAppsResponse, error) {
	op := "apps.List"
	sl.Log.Info("Listing apps", slog.String("op", op))

	resp, err := s.client.Apps(ctx, req)
	if err != nil {
		sl.Log.Error("List failed", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	return resp, nil
}
