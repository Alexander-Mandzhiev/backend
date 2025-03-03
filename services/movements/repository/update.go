package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/movements"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, movement *movements.MovementResponse) error {
	op := "repository.Update"
	sl.Log.Debug("Updating movement", slog.Int("id", int(movement.Id)), slog.String("op", op))
	query := `UPDATE movements SET product_id = ?, from_location_id = ?, to_location_id = ?, user_id = ?, comment = ? WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, movement.ProductId, movement.FromLocationId, movement.ToLocationId,
		movement.UserId, movement.Comment, movement.Id)
	if err != nil {
		sl.Log.Error("Failed to update movement", sl.Err(err), slog.Int("id", int(movement.Id)), slog.String("op", op))
		return fmt.Errorf("failed to update movement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sl.Log.Error("Failed to get rows affected", sl.Err(err), slog.Int64("id", movement.Id), slog.String("op", op))
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		sl.Log.Warn("Movement not found for update", slog.Int64("id", movement.Id), slog.String("op", op))
		return ErrMovementNotFound
	}

	sl.Log.Info("Movement updated successfully", slog.Int64("id", movement.Id), slog.String("op", op))
	return nil
}
