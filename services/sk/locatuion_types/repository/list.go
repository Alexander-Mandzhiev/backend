package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/location_types"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) List(ctx context.Context) ([]*location_types.LocationTypeResponse, error) {
	op := "repository.List"
	query := `SELECT id, name, description FROM location_types`

	sl.Log.Debug("Fetching all location types", slog.String("op", op))

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		sl.Log.Error("Failed to fetch location types", slog.String("op", op), slog.Any("error", err))
		return nil, fmt.Errorf("failed to fetch location types: %w", err)
	}
	defer rows.Close()

	var locTypes []*location_types.LocationTypeResponse
	for rows.Next() {
		var locType location_types.LocationTypeResponse
		if err = rows.Scan(&locType.Id, &locType.Name, &locType.Description); err != nil {
			sl.Log.Error("Failed to scan location type row", slog.String("op", op), slog.Any("error", err))
			return nil, fmt.Errorf("failed to scan location type row: %w", err)
		}
		locTypes = append(locTypes, &locType)
	}

	if err = rows.Err(); err != nil {
		sl.Log.Error("Error during fetching location types", slog.String("op", op), slog.Any("error", err))
		return nil, fmt.Errorf("error during fetching location types: %w", err)
	}

	sl.Log.Info("Location types fetched successfully", slog.String("op", op), slog.Int("count", len(locTypes)))
	return locTypes, nil
}
