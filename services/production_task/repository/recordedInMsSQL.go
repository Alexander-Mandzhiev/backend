package repository

import "context"

func (r *Repository) RecordedInMsSQL(ctx context.Context, ids []int) error {
	// Моковая реализация записи в MSSQL через Firebird
	if len(ids) == 0 {
		return nil
	}
	// Здесь можно добавить логику проверки ID
	return nil
}
