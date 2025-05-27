package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/movements"
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (r *Repository) Movements(ctx context.Context, productID int64) ([]*movements.MovementResponse, error) {
	op := "repository.MovementsByProduct"
	sl.Log.Debug("Fetching movements for product", slog.String("op", op), slog.Int64("product_id", productID))
	query := `SELECT id, product_id, from_location_id, to_location_id, user_id, comment, created_at, removed_at FROM movements WHERE product_id = ?`

	rows, err := r.db.QueryContext(ctx, query, productID)
	if err != nil {
		sl.Log.Error("Failed to fetch movements", sl.Err(err), slog.String("op", op))
		return nil, fmt.Errorf("failed to fetch movements: %w", err)
	}
	defer rows.Close()

	var data []*movements.MovementResponse

	for rows.Next() {
		var movement movements.MovementResponse
		var createdAt, removedAt sql.NullTime

		if err = rows.Scan(&movement.Id, &movement.ProductId, &movement.FromLocationId, &movement.ToLocationId,
			&movement.UserId, &movement.Comment, &createdAt, &removedAt); err != nil {
			sl.Log.Error("Failed to scan movement row", sl.Err(err), slog.String("op", op))
			return nil, fmt.Errorf("failed to scan movement row: %w", err)
		}

		if createdAt.Valid {
			movement.CreatedAt = timestamppb.New(createdAt.Time)
		}
		if removedAt.Valid {
			movement.RemovedAt = timestamppb.New(removedAt.Time)
		}

		data = append(data, &movement)
	}

	if err = rows.Err(); err != nil {
		sl.Log.Error("Error occurred while iterating over rows", sl.Err(err), slog.String("op", op))
		return nil, fmt.Errorf("error occurred while iterating over rows: %w", err)
	}

	sl.Log.Info("Movements fetched successfully", slog.Int("count", len(data)), slog.String("op", op))
	return data, nil
}
