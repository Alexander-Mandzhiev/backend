package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/apps"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Apps(ctx context.Context) ([]*app_provider.App, error) {
	const op = "storage.mssql.Apps"
	query := "SELECT id, name, secret FROM apps"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		sl.Log.Error(op, slog.String("message", "Failed to fetch apps"), slog.Any("error", err))
		return nil, fmt.Errorf("%s: failed to fetch apps: %w", op, err)
	}
	defer rows.Close()

	var apps []*app_provider.App
	for rows.Next() {
		var a app_provider.App
		if err = rows.Scan(&a.Id, &a.Name, &a.Secret); err != nil {
			sl.Log.Error(op, slog.String("message", "Failed to scan app"), slog.Any("error", err))
			return nil, fmt.Errorf("%s: failed to scan app: %w", op, err)
		}
		apps = append(apps, &a)
	}

	if err = rows.Err(); err != nil {
		sl.Log.Error(op, slog.String("message", "Error iterating over rows"), slog.Any("error", err))
		return nil, fmt.Errorf("%s: error iterating over rows: %w", op, err)
	}

	sl.Log.Info(op, slog.String("message", "Apps fetched successfully"), slog.Int("count", len(apps)))
	return apps, nil
}
