package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/movements"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (r *Repository) Movement(ctx context.Context, id int64) (*movements.MovementResponse, error) {
	op := "repository.Movement"
	sl.Log.Debug("Fetching movement by ID", slog.Int64("id", id), slog.String("op", op))

	query := `SELECT id, product_id, from_location_id, to_location_id, user_id, comment, created_at, removed_at FROM movements WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, id)
	var movement movements.MovementResponse
	var createdAt, removedAt sql.NullTime

	err := row.Scan(&movement.Id, &movement.ProductId, &movement.FromLocationId, &movement.ToLocationId,
		&movement.UserId, &movement.Comment, &createdAt, &removedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			sl.Log.Warn("Movement not found", slog.Int64("id", id), slog.String("op", op))
			return nil, ErrMovementNotFound
		}
		sl.Log.Error("Failed to fetch movement", sl.Err(err), slog.Int64("id", id), slog.String("op", op))
		return nil, fmt.Errorf("failed to fetch movement: %w", err)
	}

	if createdAt.Valid {
		movement.CreatedAt = timestamppb.New(createdAt.Time)
	}
	if removedAt.Valid {
		movement.RemovedAt = timestamppb.New(removedAt.Time)
	}

	sl.Log.Info("Movement fetched successfully", slog.Int64("id", movement.Id), slog.String("op", op))
	return &movement, nil
}
