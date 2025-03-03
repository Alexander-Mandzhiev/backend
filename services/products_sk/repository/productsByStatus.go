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

func (r *Repository) ProductsByStatus(ctx context.Context, statusID int32) ([]*products_sk.ProductResponse, error) {
	op := "repository.ProductsByStatus"
	sl.Log.Debug("Fetching products by active status", slog.Int64("status_id", int64(statusID)), slog.String("op", op))

	query := `SELECT p.id, p.part_name, p.nomenclature, p.number_frame, p.weight_sp_kg, p.weight_gp_kg, p.manufacturing_date
        FROM products_sk p
        INNER JOIN products_sk_statuses ps ON p.id = ps.product_id
        WHERE ps.status_id = ? AND ps.active = 1`

	rows, err := r.db.QueryContext(ctx, query, statusID)
	if err != nil {
		sl.Log.Error("Failed to fetch products by status", sl.Err(err), slog.String("op", op))
		return nil, fmt.Errorf("failed to fetch products by status: %w", err)
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
