package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

var (
	ErrAppNotFound = errors.New("app not found")
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
