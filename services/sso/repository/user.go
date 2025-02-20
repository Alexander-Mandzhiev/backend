package repository

import (
	"backend/services/sso/models"
	"context"
	"database/sql"
	"fmt"
)

func (r *Repository) User(ctx context.Context, password int) (models.User, error) {
	const op = "repository.User"
	var user models.User
	query := `SELECT USR_ID, USR_NAME, USR_PASS, USR_FLAGMASTER, USR_PREF, USR_ACTIVE, USR_CREATEDT, USR_UPDDT FROM USR WHERE USR_PASS = ?`

	row := r.db.QueryRow(query, password)
	err := row.Scan(&user.Id, &user.Name, &user.Password, &user.FlagMaster, &user.Prefix, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("%s. user not found: %w", op, err)
		}
		return user, fmt.Errorf("%s. internal server error: %w", op, err)
	}
	return user, nil
}
