package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/statuses"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, status *statuses.StatusResponse) error {
	op := "repository.Update"
	sl.Log.Debug("Updating status", slog.Int("id", int(status.Id)), slog.String("op", op))
	query := `UPDATE statuses SET name = ?, description = ? WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, status.Name, status.Description, status.Id)
	if err != nil {
		sl.Log.Error("Failed to update status", sl.Err(err), slog.Int("id", int(status.Id)), slog.String("op", op))
		return fmt.Errorf("failed to update status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sl.Log.Error("Failed to get rows affected", sl.Err(err), slog.Int("id", int(status.Id)), slog.String("op", op))
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		sl.Log.Warn("Status not found for update", slog.Int("id", int(status.Id)), slog.String("op", op))
		return ErrStatusNotFound
	}

	sl.Log.Info("Status updated successfully", slog.Int("id", int(status.Id)), slog.String("op", op))
	return nil
}
