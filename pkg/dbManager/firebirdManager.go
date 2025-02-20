package dbManager

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	sl "backend/pkg/logger"
	_ "github.com/nakagami/firebirdsql"
)

func OpenFirebirdConnection(storagePath string, maxOpenConnections, maxIdleConnections int, connectionMaxLifetime time.Duration) (*sql.DB, error) {
	const op = "firebird_manager.OpenFirebirdConnection"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sql.Open("firebirdsql", storagePath)
	if err != nil {
		sl.Log.Error(op, slog.String("message", "Error opening connection to Firebird database"), slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetConnMaxLifetime(connectionMaxLifetime)

	if err = db.PingContext(ctx); err != nil {
		sl.Log.Error(op, slog.String("message", "Error testing Firebird database connection"), slog.Any("error", err))
		if closeErr := db.Close(); closeErr != nil {
			sl.Log.Warn(op, slog.String("message", "Failed to close database connection during cleanup"), slog.Any("error", closeErr))
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	sl.Log.Info(op, slog.String("message", "Opened connection to Firebird database"))
	return db, nil
}

func CloseFirebirdConnection(db *sql.DB) error {
	const op = "firebird_manager.CloseFirebirdConnection"
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
			sl.Log.Error(op, slog.String("message", "Error closing Firebird database connection"), slog.Any("error", err))
			return fmt.Errorf("%s: %w", op, err)
		}
		sl.Log.Info(op, slog.String("message", "Firebird database connection closed"))
	case <-ctx.Done():
		sl.Log.Warn(op, slog.String("message", "Timeout while closing Firebird database connection"), slog.Any("error", ctx.Err()))
		return fmt.Errorf("%s: timeout while closing database connection", op)
	}

	return nil
}
