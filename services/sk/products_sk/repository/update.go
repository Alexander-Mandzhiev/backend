package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/products_sk"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, product *products_sk.ProductResponse) error {
	op := "repository.Update"
	sl.Log.Debug("Updating product", slog.Int64("id", product.Id), slog.String("op", op))
	query := `UPDATE products_sk SET part_name = ?, nomenclature = ?, number_frame = ?, weight_sp_kg = ?, weight_gp_kg = ?, manufacturing_date = ? WHERE id = ?`
	manufacturingDate := product.ManufacturingDate.AsTime()
	result, err := r.db.ExecContext(ctx, query, product.PartName, product.Nomenclature, product.NumberFrame, product.WeightSpKg, product.WeightGpKg, manufacturingDate, product.Id)
	if err != nil {
		sl.Log.Error("Failed to update product", sl.Err(err), slog.Int64("id", product.Id), slog.String("op", op))
		return fmt.Errorf("failed to update product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sl.Log.Error("Failed to get rows affected", sl.Err(err), slog.Int64("id", product.Id), slog.String("op", op))
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		sl.Log.Warn("Product not found for update", slog.Int64("id", product.Id), slog.String("op", op))
		return ErrProductSkNotFound
	}

	sl.Log.Info("Product updated successfully", slog.Int64("id", product.Id), slog.String("op", op))
	return nil
}
