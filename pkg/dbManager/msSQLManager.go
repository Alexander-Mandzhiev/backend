package dbManager

import (
	sl "backend/pkg/logger"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"
)

func OpenMSSQLConnection(storagePath string, maxOpenConnections, maxIdleConnections int, connectionMaxLifetime time.Duration) (*sql.DB, error) {
	const op = "db_manager.OpenMSSQLConnection"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sql.Open("mssql", storagePath)
	if err != nil {
		sl.Log.Error(op, slog.String("message", "Error opening connection to MSSQL database"), slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetConnMaxLifetime(connectionMaxLifetime)

	err = db.PingContext(ctx)
	if err != nil {
		sl.Log.Error(op, slog.String("message", "Error testing MSSQL database connection"), slog.Any("error", err))
		_ = db.Close()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	sl.Log.Info(op, slog.String("message", "Opened connection to MSSQL database"))
	return db, nil
}

func CloseMSSQLConnection(db *sql.DB) error {
	const op = "db_manager.CloseMSSQLConnection"

	if db == nil {
		sl.Log.Warn(op, slog.String("message", "Database connection is already closed or not initialized"))
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	errCh := make(chan error, 1)
	go func() {
		errCh <- db.Close()
	}()

	select {
	case err := <-errCh:
		if err != nil {
			sl.Log.Error(op, slog.String("message", "Error closing MSSQL database connection"), slog.Any("error", err))
			return fmt.Errorf("%s: %w", op, err)
		}
		sl.Log.Info(op, slog.String("message", "MSSQL database connection closed"))
	case <-ctx.Done():
		sl.Log.Warn(op, slog.String("message", "Timeout while closing MSSQL database connection"), slog.Any("error", ctx.Err()))
		return fmt.Errorf("%s: timeout while closing database connection", op)
	}

	return nil
}
