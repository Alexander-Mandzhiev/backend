package repository

import (
	sl "backend/pkg/logger"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id int32) error {
	const op = "storage.mssql.DeleteApp"
	query := "DELETE FROM apps WHERE id = ?"

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		sl.Log.Error(op, slog.String("message", "Failed to delete app"), slog.Int("app_id", int(id)), slog.Any("error", err))
		return fmt.Errorf("%s: failed to delete app: %w", op, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		sl.Log.Warn(op, slog.String("message", "App not found"), slog.Int("app_id", int(id)))
		return fmt.Errorf("%s: %w", op, ErrAppNotFound)
	}

	sl.Log.Info(op, slog.String("message", "App deleted successfully"), slog.Int("app_id", int(id)))
	return nil
}
