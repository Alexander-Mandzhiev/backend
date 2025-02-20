package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/apps"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, app *app_provider.App) error {
	const op = "storage.mssql.UpdateApp"
	query := "UPDATE apps SET name = ?, secret = ? WHERE id = ?"

	result, err := r.db.ExecContext(ctx, query, app.Name, app.Secret, app.Id)
	if err != nil {
		sl.Log.Error(op, slog.String("message", "Failed to update app"), slog.Int("app_id", int(app.Id)), slog.Any("error", err))
		return fmt.Errorf("%s: failed to update app: %w", op, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		sl.Log.Warn(op, slog.String("message", "App not found"), slog.Int("app_id", int(app.Id)))
		return fmt.Errorf("%s: %w", op, ErrAppNotFound)
	}

	sl.Log.Info(op, slog.String("message", "App updated successfully"), slog.Int("app_id", int(app.Id)))
	return nil
}
