package repository

import (
	sl "backend/pkg/logger"
	"backend/services/location/service"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, productID int64, statusID int32) error {
	op := "repository.Delete"
	sl.Log.Debug("Deactivating product status", slog.String("op", op), slog.Int64("product_id", productID), slog.Int64("status_id", int64(statusID)))
	query := `UPDATE products_sk_statuses SET active = 0 WHERE product_id = ? AND status_id = ?`
	result, err := r.db.ExecContext(ctx, query, productID, statusID)
	if err != nil {
		sl.Log.Error("Failed to deactivate product status", sl.Err(err), slog.String("op", op))
		return fmt.Errorf("failed to deactivate product status: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sl.Log.Error("Failed to get rows affected", sl.Err(err), slog.String("op", op))
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		sl.Log.Warn("Product status not found for deactivation", slog.Int64("product_id", productID), slog.Int64("status_id", int64(statusID)), slog.String("op", op))
		return service.ErrLocationNotFound
	}

	sl.Log.Info("Product status deactivated successfully", slog.Int64("product_id", productID), slog.Int64("status_id", int64(statusID)), slog.String("op", op))
	return nil
}
