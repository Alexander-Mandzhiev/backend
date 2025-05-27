package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/location_types"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

var (
	ErrNotFound = fmt.Errorf("not found")
)

func (r *Repository) Update(ctx context.Context, req *location_types.UpdateLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	const op = "repository.Update"
	const query = `
        UPDATE location_types SET name = ?, description = ?
        OUTPUT INSERTED.id, INSERTED.name, INSERTED.description WHERE id = ?`

	sl.Log.Debug("Executing update query", slog.String("op", op), slog.Int("id", int(req.GetId())))

	var response location_types.LocationTypeResponse
	err := r.db.QueryRowContext(ctx, query, req.GetName(), req.GetDescription(), req.GetId()).Scan(&response.Id, &response.Name, &response.Description)

	switch {
	case err == sql.ErrNoRows:
		sl.Log.Warn("Location type not found", slog.String("op", op), slog.Int("id", int(req.GetId())))
		return nil, ErrNotFound
	case err != nil:
		sl.Log.Error("Update failed", slog.String("op", op), slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	sl.Log.Info("Location type updated", slog.String("op", op), slog.Int("id", int(response.GetId())))
	return &response, nil
}
