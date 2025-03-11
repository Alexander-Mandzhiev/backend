package repository

import (
	sl "backend/pkg/logger"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) RecordedOutMsSQL(ctx context.Context, ids []int64) error {
	if len(ids) == 0 {
		return nil
	}

	query := "EXECUTE BLOCK AS BEGIN\n"
	for _, id := range ids {
		query += fmt.Sprintf("DELETE FROM reports_recorded_mssql WHERE OPERFAS_ID = %d;\n", id)
	}
	query += "END;"

	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		sl.Log.Error("Failed to delete records", slog.String("error", err.Error()))
		return fmt.Errorf("failed to delete from reports_recorded_mssql: %w", err)
	}

	return nil
}
