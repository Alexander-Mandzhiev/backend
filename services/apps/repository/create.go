package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/apps"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, app *app_provider.App) (int32, error) {
	const op = "storage.mssql.CreateApp"
	query := `INSERT INTO apps (name, secret) OUTPUT INSERTED.id VALUES (?, ?)`
	var id int32

	row := r.db.QueryRowContext(ctx, query, app.Name, app.Secret)
	if err := row.Scan(&id); err != nil {
		sl.Log.Error(op, slog.String("message", "Failed to create app"), slog.String("app_name", app.Name), slog.Any("error", err))
		return 0, fmt.Errorf("%s: failed to create app: %w", op, err)
	}

	sl.Log.Info(op, slog.String("message", "App created successfully"), slog.Int("app_id", int(id)))
	return id, nil
}
