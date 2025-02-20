package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/nakagami/firebirdsql"
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
		return nil, fmt.Errorf("database connection is nil")
	}
	return &Repository{db: db}, nil
}
