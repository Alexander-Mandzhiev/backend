package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/product_status_history"
	"backend/services/sk/product_status_history/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (r *Repository) GetByID(ctx context.Context, id int64) (*product_status_history.ProductSkStatusResponse, error) {
	const op = "repository.GetByID"
	logger := sl.Log.With(slog.String("op", op))

	query := `
        SELECT 
            id, 
            product_id, 
            status_id, 
            created_at 
        FROM product_status_history 
        WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, id)

	var (
		dbID      int64
		productID int64
		statusID  int32
		createdAt sql.NullTime
	)

	err := row.Scan(&dbID, &productID, &statusID, &createdAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Warn("status not found", slog.Int64("id", id))
			return nil, service.ErrProductStatusNotFound
		}
		logger.Error("failed to get status", sl.Err(err), slog.Int64("id", id))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	response := &product_status_history.ProductSkStatusResponse{
		Id:        dbID,
		ProductId: productID,
		StatusId:  statusID,
	}

	if createdAt.Valid {
		response.CreatedAt = timestamppb.New(createdAt.Time)
	}

	logger.Debug("status retrieved", slog.Int64("id", id))
	return response, nil
}
