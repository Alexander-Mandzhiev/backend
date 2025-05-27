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

func (r *Repository) Create(ctx context.Context, req *product_status_history.CreateStatusRequest) (*product_status_history.ProductSkStatusResponse, error) {
	const op = "repository.Create"
	logger := sl.Log.With(slog.String("op", op))

	if req == nil {
		logger.Error("nil request")
		return nil, fmt.Errorf("request is nil")
	}

	query := `
        INSERT INTO product_status_history (product_id, status_id) 
        OUTPUT INSERTED.id, INSERTED.created_at
        VALUES (?, ?)`

	var (
		id        int64
		createdAt sql.NullTime
	)

	err := r.db.QueryRowContext(ctx, query, req.ProductId, req.StatusId).Scan(&id, &createdAt)

	if err != nil {
		logger.Error("failed to create status", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	response := &product_status_history.ProductSkStatusResponse{
		Id:        id,
		ProductId: req.ProductId,
		StatusId:  req.StatusId,
	}

	if createdAt.Valid {
		response.CreatedAt = timestamppb.New(createdAt.Time)
	}

	logger.Info("status created", slog.Int64("id", id))
	return response, nil
}
