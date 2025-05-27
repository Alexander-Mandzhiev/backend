package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/locations"
	"backend/services/sk/location/service"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

func (r *Repository) Location(ctx context.Context, id int) (*locations.UpdateLocationRequest, error) {
	op := "repository.Location"
	query := `SELECT id, name, type_id, capacity, current_load FROM locations WHERE id = ?`

	sl.Log.Debug("Fetching location by ID", slog.String("op", op), slog.Int("id", id))

	var loc locations.UpdateLocationRequest
	err := r.db.QueryRowContext(ctx, query, id).Scan(&loc.Id, &loc.Name, &loc.TypeId, &loc.Capacity, &loc.CurrentLoad)
	if err == sql.ErrNoRows {
		sl.Log.Warn("Location not found", slog.String("op", op), slog.Int("id", id))
		return nil, service.ErrLocationNotFound
	} else if err != nil {
		sl.Log.Error("Failed to fetch location", slog.String("op", op), slog.Any("error", err), slog.Int("id", id))
		return nil, fmt.Errorf("failed to fetch location: %w", err)
	}

	sl.Log.Info("Location fetched successfully", slog.String("op", op), slog.Int("id", int(loc.Id)))
	return &loc, nil
}
