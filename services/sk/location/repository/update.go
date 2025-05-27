package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/locations"
	"backend/services/sk/location/service"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, location *locations.UpdateLocationRequest) error {
	op := "repository.Update"
	query := `UPDATE locations SET name = ?, type_id = ?, capacity = ?, current_load = ? WHERE id = ?`
	sl.Log.Debug("Updating location", slog.String("op", op), slog.Any("location", location))
	result, err := r.db.ExecContext(ctx, query, location.GetName(), location.GetTypeId(), location.GetCapacity(), location.GetCurrentLoad(), location.GetId())
	if err != nil {
		sl.Log.Error("Failed to update location", slog.String("op", op), slog.Any("error", err), slog.Int("id", int(location.GetId())))
		return fmt.Errorf("failed to update location: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sl.Log.Error("Failed to check rows affected", slog.String("op", op), slog.Any("error", err), slog.Int("id", int(location.GetId())))
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		sl.Log.Warn("Location not found during update", slog.String("op", op), slog.Int("id", int(location.GetId())))
		return service.ErrLocationNotFound
	}

	sl.Log.Info("Location updated successfully", slog.String("op", op), slog.Int("id", int(location.GetId())))
	return nil
}
