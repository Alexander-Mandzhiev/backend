package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/statuses"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Statuses(ctx context.Context) ([]*statuses.StatusResponse, error) {
	op := "repository.Statuses"
	sl.Log.Debug("Fetching all statuses", slog.String("op", op))
	query := `SELECT id, name, description FROM statuses`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		sl.Log.Error("Failed to fetch statuses", sl.Err(err), slog.String("op", op))
		return nil, fmt.Errorf("failed to fetch statuses: %w", err)
	}
	defer rows.Close()

	var statusesList []*statuses.StatusResponse
	for rows.Next() {
		var status statuses.StatusResponse
		if err = rows.Scan(&status.Id, &status.Name, &status.Description); err != nil {
			sl.Log.Error("Failed to scan status row", sl.Err(err), slog.String("op", op))
			return nil, fmt.Errorf("failed to scan status row: %w", err)
		}
		statusesList = append(statusesList, &status)
	}

	if err = rows.Err(); err != nil {
		sl.Log.Error("Error occurred while iterating over rows", sl.Err(err), slog.String("op", op))
		return nil, fmt.Errorf("error occurred while iterating over rows: %w", err)
	}

	sl.Log.Info("Statuses fetched successfully", slog.Int("count", len(statusesList)), slog.String("op", op))
	return statusesList, nil
}
