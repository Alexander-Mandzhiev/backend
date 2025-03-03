package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk"
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (r *Repository) Products(ctx context.Context) ([]*products_sk.ProductResponse, error) {
	op := "repository.List"
	sl.Log.Debug("Fetching all products", slog.String("op", op))

	query := `
        SELECT id, part_name, nomenclature, number_frame, weight_sp_kg, weight_gp_kg, manufacturing_date
        FROM products_sk
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		sl.Log.Error("Failed to fetch products", sl.Err(err), slog.String("op", op))
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}
	defer rows.Close()

	var products []*products_sk.ProductResponse
	for rows.Next() {
		var product products_sk.ProductResponse
		var manufacturingDate sql.NullTime

		if err = rows.Scan(&product.Id, &product.PartName, &product.Nomenclature, &product.NumberFrame, &product.WeightSpKg, &product.WeightGpKg, &manufacturingDate); err != nil {
			sl.Log.Error("Failed to scan product row", sl.Err(err), slog.String("op", op))
			return nil, fmt.Errorf("failed to scan product row: %w", err)
		}

		if manufacturingDate.Valid {
			product.ManufacturingDate = timestamppb.New(manufacturingDate.Time)
		}

		products = append(products, &product)
	}

	if err = rows.Err(); err != nil {
		sl.Log.Error("Error occurred while iterating over rows", sl.Err(err), slog.String("op", op))
		return nil, fmt.Errorf("error occurred while iterating over rows: %w", err)
	}

	sl.Log.Info("Products fetched successfully", slog.Int("count", len(products)), slog.String("op", op))
	return products, nil
}
