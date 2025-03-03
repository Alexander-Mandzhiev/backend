package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, product *products_sk.ProductResponse) error {
	op := "repository.Create"
	sl.Log.Debug("Creating new product", slog.String("op", op), slog.Int64("id", product.Id))
	query := `INSERT INTO products_sk (id, part_name, nomenclature, number_frame, weight_sp_kg, weight_gp_kg, manufacturing_date) VALUES (?, ?, ?, ?, ?, ?, ?)`
	manufacturingDate := product.ManufacturingDate.AsTime()
	result, err := r.db.ExecContext(ctx, query, product.Id, product.PartName, product.Nomenclature, product.NumberFrame, product.WeightSpKg, product.WeightGpKg, manufacturingDate)
	if err != nil {
		sl.Log.Error("Failed to create product", sl.Err(err), slog.String("op", op))
		return fmt.Errorf("failed to create product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sl.Log.Error("Failed to get rows affected", sl.Err(err), slog.String("op", op))
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		sl.Log.Warn("Product not created", slog.Int64("id", product.Id), slog.String("op", op))
		return fmt.Errorf("product with ID %d not created", product.Id)
	}

	sl.Log.Info("Product created successfully", slog.Int64("id", product.Id), slog.String("op", op))
	return nil
}
