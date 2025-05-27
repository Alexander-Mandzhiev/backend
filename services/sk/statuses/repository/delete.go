package repository

import (
	sl "backend/pkg/logger"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id int) error {
	op := "repository.Delete"
	sl.Log.Debug("Deleting status", slog.Int("id", id), slog.String("op", op))
	query := `DELETE FROM statuses WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		sl.Log.Error("Failed to delete status", sl.Err(err), slog.Int("id", id), slog.String("op", op))
		return fmt.Errorf("failed to delete status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sl.Log.Error("Failed to get rows affected", sl.Err(err), slog.Int("id", id), slog.String("op", op))
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		sl.Log.Warn("Status not found for deletion", slog.Int("id", id), slog.String("op", op))
		return ErrStatusNotFound
	}

	sl.Log.Info("Status deleted successfully", slog.Int("id", id), slog.String("op", op))
	return nil
}
