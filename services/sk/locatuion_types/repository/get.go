package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/location_types"
	"backend/services/sk/locatuion_types/service"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

func (r *Repository) Get(ctx context.Context, id int) (*location_types.LocationTypeResponse, error) {
	op := "repository.Get"
	query := `SELECT id, name, description FROM location_types WHERE id = ?`

	sl.Log.Debug("Fetching location type by ID", slog.String("op", op), slog.Int("id", id))

	var locType location_types.LocationTypeResponse
	err := r.db.QueryRowContext(ctx, query, id).Scan(&locType.Id, &locType.Name, &locType.Description)
	if err == sql.ErrNoRows {
		sl.Log.Warn("Location type not found", slog.String("op", op), slog.Int("id", id))
		return nil, service.ErrLocationTypesNotFound
	} else if err != nil {
		sl.Log.Error("Failed to fetch location type", slog.String("op", op), slog.Any("error", err), slog.Int("id", id))
		return nil, fmt.Errorf("failed to fetch location type: %w", err)
	}

	sl.Log.Info("Location type fetched successfully", slog.String("op", op), slog.Int("id", int(locType.Id)))
	return &locType, nil
}
