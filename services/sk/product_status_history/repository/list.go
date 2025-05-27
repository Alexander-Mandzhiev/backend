package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/product_status_history"
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (r *Repository) ListByProduct(ctx context.Context, productID int64) ([]*product_status_history.ProductSkStatusResponse, error) {
	const op = "repository.ListByProduct"
	logger := sl.Log.With(slog.String("op", op))

	query := `
        SELECT 
            id, 
            product_id, 
            status_id, 
            created_at 
        FROM product_status_history 
        WHERE product_id = ?`

	rows, err := r.db.QueryContext(ctx, query, productID)
	if err != nil {
		logger.Error("query failed", slog.Int64("product_id", productID), sl.Err(err))
		return nil, fmt.Errorf("%s: query failed: %w", op, err)
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			logger.Warn("failed to close rows", sl.Err(closeErr))
		}
	}()

	var statuses []*product_status_history.ProductSkStatusResponse

	for rows.Next() {
		var (
			id        int64
			dbProduct int64
			statusID  int32
			createdAt sql.NullTime
		)

		if err = rows.Scan(&id, &dbProduct, &statusID, &createdAt); err != nil {
			logger.Error("row scan failed", slog.Int64("current_id", id), sl.Err(err))
			return nil, fmt.Errorf("%s: scan failed: %w", op, err)
		}

		response := &product_status_history.ProductSkStatusResponse{
			Id:        id,
			ProductId: dbProduct,
			StatusId:  statusID,
		}

		if createdAt.Valid {
			response.CreatedAt = timestamppb.New(createdAt.Time)
		} else {
			logger.Warn("null timestamp", slog.Int64("record_id", id))
		}

		statuses = append(statuses, response)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", sl.Err(err))
		return nil, fmt.Errorf("%s: iteration error: %w", op, err)
	}

	logger.Debug("results fetched", slog.Int64("product_id", productID), slog.Int("count", len(statuses)))
	return statuses, nil
}
