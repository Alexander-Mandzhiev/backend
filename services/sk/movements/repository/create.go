package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/movements"
	"context"
	"fmt"
	"log/slog"
	"time"
)

func (r *Repository) Create(ctx context.Context, movement *movements.MovementResponse) (int64, error) {
	op := "repository.Create"
	sl.Log.Debug("Creating new movement", slog.String("op", op))
	query := `INSERT INTO movements (product_id, from_location_id, to_location_id, user_id, comment, created_at) OUTPUT INSERTED.id VALUES (?, ?, ?, ?, ?, ?)`
	var id int64
	now := time.Now()
	err := r.db.QueryRowContext(ctx, query, movement.ProductId, movement.FromLocationId, movement.ToLocationId, movement.UserId, movement.Comment, now).Scan(&id)
	if err != nil {
		sl.Log.Error("Failed to create movement", sl.Err(err), slog.String("op", op))
		return 0, fmt.Errorf("failed to create movement: %w", err)
	}

	sl.Log.Info("Movement created successfully", slog.Int64("id", id), slog.String("op", op))
	return id, nil
}
