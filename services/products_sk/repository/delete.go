package repository

import (
	sl "backend/pkg/logger"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id int64) error {
	op := "repository.Delete"
	sl.Log.Debug("Deleting product", slog.Int64("id", id), slog.String("op", op))
	query := `DELETE FROM products_sk WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		sl.Log.Error("Failed to delete product", sl.Err(err), slog.Int64("id", id), slog.String("op", op))
		return fmt.Errorf("failed to delete product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sl.Log.Error("Failed to get rows affected", sl.Err(err), slog.Int64("id", id), slog.String("op", op))
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		sl.Log.Warn("Product not found for deletion", slog.Int64("id", id), slog.String("op", op))
		return ErrProductSkNotFound
	}

	sl.Log.Info("Product deleted successfully", slog.Int64("id", id), slog.String("op", op))
	return nil
}
