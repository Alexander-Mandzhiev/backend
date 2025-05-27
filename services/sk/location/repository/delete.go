package repository

import (
	sl "backend/pkg/logger"
	"backend/services/sk/location/service"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id int) error {
	op := "repository.Delete"
	query := `DELETE FROM locations WHERE id = ?`
	sl.Log.Debug("Deleting location", slog.String("op", op), slog.Int("id", id))

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		sl.Log.Error("Failed to delete location", slog.String("op", op), slog.Any("error", err), slog.Int("id", id))
		return fmt.Errorf("failed to delete location: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sl.Log.Error("Failed to check rows affected", slog.String("op", op), slog.Any("error", err), slog.Int("id", id))
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		sl.Log.Warn("Location not found during delete", slog.String("op", op), slog.Int("id", id))
		return service.ErrLocationNotFound
	}

	sl.Log.Info("Location deleted successfully", slog.String("op", op), slog.Int("id", id))
	return nil
}
