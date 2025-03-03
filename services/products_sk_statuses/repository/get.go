package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk_statuses"
	"backend/services/products_sk_statuses/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) Get(ctx context.Context, productID int64, statusID int32) (*products_sk_statuses.ProductSkStatusResponse, error) {
	op := "repository.Get"
	sl.Log.Debug("Fetching product status", slog.String("op", op), slog.Int64("product_id", productID), slog.Int64("status_id", int64(statusID)))
	query := `SELECT product_id, status_id, active, created_at FROM products_sk_statuses WHERE product_id = ? AND status_id = ?`
	row := r.db.QueryRowContext(ctx, query, productID, statusID)
	var status products_sk_statuses.ProductSkStatusResponse
	var createdAt time.Time
	err := row.Scan(&status.ProductId, &status.StatusId, &status.Active, &createdAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			sl.Log.Warn("Product status not found", slog.Int64("product_id", productID), slog.Int64("status_id", int64(statusID)), slog.String("op", op))
			return nil, service.ErrProductStatusNotFound
		}
		sl.Log.Error("Failed to fetch product status", sl.Err(err), slog.String("op", op))
		return nil, fmt.Errorf("failed to fetch product status: %w", err)
	}

	status.CreatedAt = timestamppb.New(createdAt)
	return &status, nil
}
