package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (r *Repository) Product(ctx context.Context, id int64) (*products_sk.ProductResponse, error) {
	op := "repository.Product"
	sl.Log.Debug("Fetching product by ID", slog.Int64("id", id), slog.String("op", op))
	query := `SELECT id, part_name, nomenclature, number_frame, weight_sp_kg, weight_gp_kg, manufacturing_date FROM products_sk WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)
	var product products_sk.ProductResponse
	var manufacturingDate sql.NullTime

	if err := row.Scan(&product.Id, &product.PartName, &product.Nomenclature, &product.NumberFrame, &product.WeightSpKg, &product.WeightGpKg, &manufacturingDate); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			sl.Log.Warn("Product not found", slog.Int64("id", id), slog.String("op", op))
			return nil, ErrProductSkNotFound
		}
		sl.Log.Error("Failed to fetch product", sl.Err(err), slog.Int64("id", id), slog.String("op", op))
		return nil, fmt.Errorf("failed to fetch product: %w", err)
	}

	if manufacturingDate.Valid {
		product.ManufacturingDate = timestamppb.New(manufacturingDate.Time)
	}

	sl.Log.Info("Product fetched successfully", slog.Int64("id", product.Id), slog.String("op", op))
	return &product, nil
}
