package repository

import (
	sl "backend/pkg/logger"
	"backend/services/sk/product_status_history/service"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id int64) error {
	const op = "repository.Delete"
	logger := sl.Log.With(slog.String("op", op))

	if id <= 0 {
		logger.Error("invalid ID", slog.Int64("id", id))
		return fmt.Errorf("%s: invalid ID: %w", op, service.ErrInvalidRequest)
	}

	query := `DELETE FROM product_status_history WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		logger.Error("failed to delete status", sl.Err(err), slog.Int64("id", id))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("failed to check rows affected", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		logger.Warn("status not found for deletion", slog.Int64("id", id))
		return service.ErrProductStatusNotFound
	}

	logger.Info("status deleted successfully", slog.Int64("id", id))
	return nil
}
