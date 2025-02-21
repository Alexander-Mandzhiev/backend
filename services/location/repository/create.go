package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, loc *locations.LocationResponse) (int32, error) {
	op := "repository.Create"
	query := `
        INSERT INTO locations (name, type, capacity, current_load)
        OUTPUT INSERTED.id
        VALUES (?, ?, ?, ?)
    `

	sl.Log.Debug("Creating new location", slog.String("op", op), slog.Any("location", loc))

	var id int32
	err := r.db.QueryRowContext(ctx, query, loc.GetName(), loc.GetType(),
		loc.GetCapacity(), loc.GetCurrentLoad()).Scan(&id)

	if err != nil {
		sl.Log.Error("Failed to create location", slog.String("op", op), slog.Any("error", err))
		return 0, fmt.Errorf("failed to create location: %w", err)
	}

	sl.Log.Info("Location created successfully", slog.String("op", op), slog.Int("id", int(id)))
	return id, nil
}
