package repository

import (
	sl "backend/pkg/logger"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/nakagami/firebirdsql"
	"log/slog"
)

var (
	ErrInvalidCredentials = errors.New("неправильный формат данных")
	ErrInternalServerErr  = errors.New("проблема на сервере")
	ErrUserNotFound       = errors.New("пользователь не найден")
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) (*Repository, error) {
	if db == nil {
		sl.Log.Error("Database connection is nil", slog.String("op", "repository.New"))
		return nil, fmt.Errorf("database connection is nil")
	}
	sl.Log.Info("Repository initialized", slog.String("op", "repository.New"))
	return &Repository{db: db}, nil
}
