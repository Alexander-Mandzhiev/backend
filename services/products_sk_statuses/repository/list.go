package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk_statuses"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) List(ctx context.Context) ([]*products_sk_statuses.ProductSkStatusResponse, error) {
	op := "repository.List"
	sl.Log.Debug("Fetching all active product statuses", slog.String("op", op))
	query := `SELECT product_id, status_id, active, created_at FROM products_sk_statuses WHERE active = 1`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		sl.Log.Error("Failed to fetch product statuses", sl.Err(err), slog.String("op", op))
		return nil, fmt.Errorf("failed to fetch product statuses: %w", err)
	}
	defer rows.Close()

	var data []*products_sk_statuses.ProductSkStatusResponse
	for rows.Next() {
		var status products_sk_statuses.ProductSkStatusResponse
		var createdAt time.Time

		if err = rows.Scan(&status.ProductId, &status.StatusId, &status.Active, &createdAt); err != nil {
			sl.Log.Error("Failed to scan product status row", sl.Err(err), slog.String("op", op))
			return nil, fmt.Errorf("failed to scan product status row: %w", err)
		}

		status.CreatedAt = timestamppb.New(createdAt)
		data = append(data, &status)
	}

	if err = rows.Err(); err != nil {
		sl.Log.Error("Error occurred while iterating over rows", sl.Err(err), slog.String("op", op))
		return nil, fmt.Errorf("error occurred while iterating over rows: %w", err)
	}

	sl.Log.Info("Product statuses fetched successfully", slog.Int("count", len(data)), slog.String("op", op))
	return data, nil
}
