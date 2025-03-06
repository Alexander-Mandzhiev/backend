package repository

import (
	sl "backend/pkg/logger"
	"backend/services/location/service"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id int) error {
	op := "repository.Delete"
	query := `DELETE FROM location_types WHERE id = ?`

	sl.Log.Debug("Deleting location type", slog.String("op", op), slog.Int("id", id))

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		sl.Log.Error("Failed to delete location type", slog.String("op", op), slog.Any("error", err), slog.Int("id", id))
		return fmt.Errorf("failed to delete location type: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sl.Log.Error("Failed to check rows affected", slog.String("op", op), slog.Any("error", err), slog.Int("id", id))
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		sl.Log.Warn("Location type not found during delete", slog.String("op", op), slog.Int("id", id))
		return service.ErrLocationNotFound
	}

	sl.Log.Info("Location type deleted successfully", slog.String("op", op), slog.Int("id", id))
	return nil
}
