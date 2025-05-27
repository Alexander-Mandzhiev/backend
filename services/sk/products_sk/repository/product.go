package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/products_sk"
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

	query := `SELECT 
            id, part_name, nomenclature, number_frame, count_sausage_sticks,
            weight_sp_kg, weight_gp_kg, manufacturing_date, created_at, removed_at
        FROM products_sk
        WHERE id = ? AND removed_at IS NULL`

	row := r.db.QueryRowContext(ctx, query, id)
	var product products_sk.ProductResponse

	var (
		countSausage      sql.NullInt32
		weightGp          sql.NullFloat64
		manufacturingDate sql.NullTime
		createdAt         sql.NullTime
		removedAt         sql.NullTime
	)

	if err := row.Scan(&product.Id, &product.PartName, &product.Nomenclature, &product.NumberFrame, &countSausage, &product.WeightSpKg, &weightGp, &manufacturingDate, &createdAt, &removedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			sl.Log.Warn("Product not found", slog.Int64("id", id), slog.String("op", op))
			return nil, ErrProductSkNotFound
		}
		sl.Log.Error("Failed to fetch product", sl.Err(err), slog.Int64("id", id), slog.String("op", op))
		return nil, fmt.Errorf("failed to fetch product: %w", err)
	}

	if countSausage.Valid {
		product.CountSausageSticks = &countSausage.Int32
	}

	if weightGp.Valid {
		product.WeightGpKg = &weightGp.Float64
	}

	if manufacturingDate.Valid {
		product.ManufacturingDate = timestamppb.New(manufacturingDate.Time)
	}

	if createdAt.Valid {
		product.CreatedAt = timestamppb.New(createdAt.Time)
	}

	if removedAt.Valid {
		product.RemovedAt = timestamppb.New(removedAt.Time)
	}

	sl.Log.Info("Product fetched successfully", slog.Int64("id", product.Id), slog.String("op", op))
	return &product, nil
}
