package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk_statuses"
	"backend/services/products_sk_statuses/service"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, productStatus *products_sk_statuses.ProductSkStatusResponse) error {
	op := "repository.Update"
	sl.Log.Debug("Updating product status", slog.String("op", op), slog.Int64("product_id", productStatus.ProductId), slog.Int64("status_id", int64(productStatus.StatusId)))
	query := `UPDATE products_sk_statuses SET active = ? WHERE product_id = ? AND status_id = ?`
	result, err := r.db.ExecContext(ctx, query, productStatus.Active, productStatus.ProductId, productStatus.StatusId)
	if err != nil {
		sl.Log.Error("Failed to update product status", sl.Err(err), slog.String("op", op))
		return fmt.Errorf("failed to update product status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sl.Log.Error("Failed to get rows affected", sl.Err(err), slog.String("op", op))
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		sl.Log.Warn("Product status not found for update", slog.Int64("product_id", productStatus.ProductId), slog.Int64("status_id", int64(productStatus.StatusId)), slog.String("op", op))
		return service.ErrProductStatusNotFound
	}

	sl.Log.Info("Product status updated successfully", slog.Int64("product_id", productStatus.ProductId), slog.Int64("status_id", int64(productStatus.StatusId)), slog.String("op", op))
	return nil
}
