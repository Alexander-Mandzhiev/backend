package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, locType *location_types.CreateLocationTypeRequest) (int32, error) {
	op := "repository.Create"
	query := `
        INSERT INTO location_types (name, description)
        OUTPUT INSERTED.id
        VALUES (?, ?)
    `

	sl.Log.Debug("Creating new location type", slog.String("op", op), slog.Any("request", locType))

	var id int32
	err := r.db.QueryRowContext(ctx, query, locType.GetName(), locType.GetDescription()).Scan(&id)

	if err != nil {
		sl.Log.Error("Failed to create location type", slog.String("op", op), slog.Any("error", err))
		return 0, fmt.Errorf("failed to create location type: %w", err)
	}

	sl.Log.Info("Location type created successfully", slog.String("op", op), slog.Int("id", int(id)))
	return id, nil
}
