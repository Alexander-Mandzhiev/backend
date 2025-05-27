package repository

import (
	sl "backend/pkg/logger"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id int64) error {
	op := "repository.Delete"
	sl.Log.Debug("Soft-deleting product", slog.Int64("id", id), slog.String("op", op))

	var exists bool
	err := r.db.QueryRowContext(ctx, "SELECT 1 FROM products_sk WHERE id = ? AND removed_at IS NULL", id).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			sl.Log.Warn("Product not found or already deleted", slog.Int64("id", id), slog.String("op", op))
			return ErrProductSkNotFound
		}
		sl.Log.Error("Failed to check product existence", sl.Err(err), slog.Int64("id", id), slog.String("op", op))
		return fmt.Errorf("failed to check product existence: %w", err)
	}

	// Выполняем мягкое удаление
	query := `UPDATE products_sk SET removed_at = GETDATE() WHERE id = ? AND removed_at IS NULL`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		sl.Log.Error("Failed to soft-delete product", sl.Err(err), slog.Int64("id", id), slog.String("op", op))
		return fmt.Errorf("failed to soft-delete product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sl.Log.Error("Failed to get rows affected", sl.Err(err), slog.Int64("id", id), slog.String("op", op))
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		sl.Log.Warn("Product not found or already deleted", slog.Int64("id", id), slog.String("op", op))
		return ErrProductSkNotFound
	}

	sl.Log.Info("Product soft-deleted successfully", slog.Int64("id", id), slog.String("op", op))
	return nil
}
