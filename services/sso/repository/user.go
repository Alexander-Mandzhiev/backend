package repository

import (
	sl "backend/pkg/logger"
	"backend/services/sso/models"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

func (r *Repository) User(ctx context.Context, password int) (models.User, error) {
	const op = "repository.User"
	sl.Log.Info("Fetching user by password", slog.Int("password", password), slog.String("op", op))

	var user models.User
	query := `SELECT USR_ID, USR_NAME, USR_PASS, USR_FLAGMASTER, USR_PREF, USR_ACTIVE, USR_CREATEDT, USR_UPDDT FROM USR WHERE USR_PASS = ?`

	row := r.db.QueryRow(query, password)
	err := row.Scan(&user.Id, &user.Name, &user.Password, &user.FlagMaster, &user.Prefix, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			sl.Log.Warn("User not found", slog.Int("password", password), slog.String("op", op))
			return user, fmt.Errorf("%s. user not found: %w", op, err)
		}
		sl.Log.Error("Error fetching user", sl.Err(err), slog.Int("password", password), slog.String("op", op))
		return user, fmt.Errorf("%s. internal server error: %w", op, err)
	}
	sl.Log.Info("User fetched successfully", slog.Int("id", user.Id), slog.String("name", user.Name), slog.String("op", op))
	return user, nil
}
