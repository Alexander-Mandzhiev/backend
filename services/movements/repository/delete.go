package repository

import (
	sl "backend/pkg/logger"
	"context"
	"fmt"
	"log/slog"
	"time"
)

func (r *Repository) Delete(ctx context.Context, id int64) error {
	op := "repository.Delete"
	sl.Log.Debug("Deleting movement", slog.Int64("id", id), slog.String("op", op))
	query := `UPDATE movements SET removed_at = ? WHERE id = ?`
	now := time.Now()
	result, err := r.db.ExecContext(ctx, query, now, id)
	if err != nil {
		sl.Log.Error("Failed to delete movement", sl.Err(err), slog.Int64("id", id), slog.String("op", op))
		return fmt.Errorf("failed to delete movement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sl.Log.Error("Failed to get rows affected", sl.Err(err), slog.Int64("id", id), slog.String("op", op))
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		sl.Log.Warn("Movement not found for deletion", slog.Int64("id", id), slog.String("op", op))
		return ErrMovementNotFound
	}

	sl.Log.Info("Movement deleted successfully", slog.Int64("id", id), slog.String("op", op))
	return nil
}
