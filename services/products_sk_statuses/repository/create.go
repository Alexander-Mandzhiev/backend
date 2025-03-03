package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk_statuses"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, productStatus *products_sk_statuses.ProductSkStatusResponse) error {
	op := "repository.Create"
	sl.Log.Debug("Creating new product status", slog.String("op", op))
	query := `INSERT INTO products_sk_statuses (product_id, status_id, active, created_at) VALUES (?, ?, ?, GETDATE())`
	_, err := r.db.ExecContext(ctx, query, productStatus.ProductId, productStatus.StatusId, productStatus.Active)
	if err != nil {
		sl.Log.Error("Failed to create product status", sl.Err(err), slog.String("op", op))
		return fmt.Errorf("failed to create product status: %w", err)
	}

	sl.Log.Info("Product status created successfully", slog.Int64("product_id", productStatus.ProductId), slog.Int64("status_id", int64(productStatus.StatusId)), slog.String("op", op))
	return nil
}
