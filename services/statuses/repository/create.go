package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, status *statuses.StatusResponse) (int, error) {
	op := "repository.Create"
	sl.Log.Debug("Creating new status", slog.String("op", op))
	query := `INSERT INTO statuses (name, description) OUTPUT INSERTED.id VALUES (?, ?)`
	var id int
	err := r.db.QueryRowContext(ctx, query, status.Name, status.Description).Scan(&id)
	if err != nil {
		sl.Log.Error("Failed to create status", sl.Err(err), slog.String("op", op))
		return 0, fmt.Errorf("failed to create status: %w", err)
	}

	sl.Log.Info("Status created successfully", slog.Int("id", id), slog.String("op", op))
	return id, nil
}
