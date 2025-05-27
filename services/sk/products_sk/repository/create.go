package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/products_sk"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, product *products_sk.ProductResponse) error {
	op := "repository.Create"
	sl.Log.Debug("Creating new product", slog.String("op", op))

	manufacturingDate := product.ManufacturingDate.AsTime()

	var countSausage interface{}
	if product.CountSausageSticks != nil {
		countSausage = *product.CountSausageSticks
	} else {
		countSausage = nil
	}

	var weightGp interface{}
	if product.WeightGpKg != nil {
		weightGp = *product.WeightGpKg
	} else {
		weightGp = nil
	}

	query := `
        INSERT INTO products_sk 
        (part_name, nomenclature, number_frame, count_sausage_sticks, weight_sp_kg, weight_gp_kg, manufacturing_date) 
        VALUES (?, ?, ?, ?, ?, ?, ?);
        SELECT SCOPE_IDENTITY();
    `

	var newID int64
	err := r.db.QueryRowContext(ctx, query, product.PartName, product.Nomenclature, product.NumberFrame, countSausage, product.WeightSpKg, weightGp, manufacturingDate).Scan(&newID)

	if err != nil {
		sl.Log.Error("Failed to create product", sl.Err(err), slog.String("op", op))
		return fmt.Errorf("failed to create product: %w", err)
	}

	product.Id = newID
	sl.Log.Info("Product created successfully", slog.Int64("id", newID), slog.String("op", op))
	return nil
}
