package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"backend/services/location/service"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, locType *location_types.UpdateLocationTypeRequest) error {
	op := "repository.Update"
	query := `UPDATE location_types SET name = ?, description = ? WHERE id = ?`

	sl.Log.Debug("Updating location type", slog.String("op", op), slog.Any("request", locType))

	result, err := r.db.ExecContext(ctx, query, locType.GetName(), locType.GetDescription(), locType.GetId())
	if err != nil {
		sl.Log.Error("Failed to update location type", slog.String("op", op), slog.Any("error", err), slog.Int("id", int(locType.GetId())))
		return fmt.Errorf("failed to update location type: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sl.Log.Error("Failed to check rows affected", slog.String("op", op), slog.Any("error", err), slog.Int("id", int(locType.GetId())))
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		sl.Log.Warn("Location type not found during update", slog.String("op", op), slog.Int("id", int(locType.GetId())))
		return service.ErrLocationNotFound
	}

	sl.Log.Info("Location type updated successfully", slog.String("op", op), slog.Int("id", int(locType.GetId())))
	return nil
}
