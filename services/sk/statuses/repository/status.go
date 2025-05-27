package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/statuses"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
)

func (r *Repository) Status(ctx context.Context, id int) (*statuses.StatusResponse, error) {
	op := "repository.Status"
	sl.Log.Debug("Fetching status by ID", slog.Int("id", id), slog.String("op", op))
	query := `SELECT id, name, description FROM statuses WHERE id = ?`
	var status statuses.StatusResponse
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&status.Id, &status.Name, &status.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			sl.Log.Warn("Status not found", slog.Int("id", id), slog.String("op", op))
			return nil, ErrStatusNotFound
		}
		sl.Log.Error("Failed to fetch status", sl.Err(err), slog.Int("id", id), slog.String("op", op))
		return nil, fmt.Errorf("failed to fetch status: %w", err)
	}

	sl.Log.Info("Status fetched successfully", slog.Int("id", int(status.Id)), slog.String("op", op))
	return &status, nil
}
