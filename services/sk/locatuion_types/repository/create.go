// Репозиторий
package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/location_types"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, locType *location_types.CreateLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	const op = "repository.Create"
	const query = `
        INSERT INTO location_types (name, description)
        OUTPUT INSERTED.id, INSERTED.name, INSERTED.description
        VALUES (?, ?)`

	sl.Log.Debug("Executing SQL query", slog.String("query", query), slog.String("op", op))

	var response location_types.LocationTypeResponse
	err := r.db.QueryRowContext(ctx, query, locType.GetName(), locType.GetDescription()).Scan(&response.Id, &response.Name, &response.Description)
	if err != nil {
		sl.Log.Error("Database operation failed", slog.String("op", op), slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &response, nil
}
