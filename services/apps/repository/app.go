package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/apps"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
)

func (r *Repository) App(ctx context.Context, id int32) (*app_provider.App, error) {
	const op = "storage.mssql.App"
	var app app_provider.App
	query := "SELECT id, name, secret FROM apps WHERE id = ?"
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&app.Id, &app.Name, &app.Secret); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			sl.Log.Warn(op, slog.String("message", "App not found"), slog.Int("app_id", int(id)))
			return nil, fmt.Errorf("%s: %w", op, ErrAppNotFound)
		}
		sl.Log.Error(op, slog.String("message", "Failed to fetch app"), slog.Int("app_id", int(id)), slog.Any("error", err))
		return nil, fmt.Errorf("%s: failed to fetch app: %w", op, err)
	}
	sl.Log.Info(op, slog.String("message", "App fetched successfully"), slog.Int("app_id", int(id)))
	return &app, nil
}
