package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Locations(ctx context.Context) ([]*locations.UpdateLocationRequest, error) {
	op := "repository.Locations"
	query := `SELECT id, name, type_id, capacity, current_load FROM locations`

	sl.Log.Debug("Fetching all locations", slog.String("op", op))

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		sl.Log.Error("Failed to fetch locations", slog.String("op", op), slog.Any("error", err))
		return nil, fmt.Errorf("failed to fetch locations: %w", err)
	}
	defer rows.Close()

	var locationsList []*locations.UpdateLocationRequest
	for rows.Next() {
		var loc locations.UpdateLocationRequest
		if err = rows.Scan(&loc.Id, &loc.Name, &loc.TypeId, &loc.Capacity, &loc.CurrentLoad); err != nil {
			sl.Log.Error("Failed to scan location row", slog.String("op", op), slog.Any("error", err))
			return nil, fmt.Errorf("failed to scan location row: %w", err)
		}
		locationsList = append(locationsList, &loc)
	}

	if err = rows.Err(); err != nil {
		sl.Log.Error("Error during fetching locations", slog.String("op", op), slog.Any("error", err))
		return nil, fmt.Errorf("error during fetching locations: %w", err)
	}

	sl.Log.Info("Locations fetched successfully", slog.String("op", op), slog.Int("count", len(locationsList)))
	return locationsList, nil
}
