package repository

import (
	sl "backend/pkg/logger"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) RecordedInMsSQL(ctx context.Context, ids []int64) error {
	if len(ids) == 0 {
		return nil
	}

	query := "EXECUTE BLOCK AS BEGIN\n"
	for _, id := range ids {
		query += fmt.Sprintf("INSERT INTO reports_recorded_mssql (OPERFAS_ID) VALUES (%d);\n", id)
	}
	query += "END;"

	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		sl.Log.Error("Failed to insert records", slog.String("error", err.Error()))
		return fmt.Errorf("failed to insert into reports_recorded_mssql: %w", err)
	}

	return nil
}
