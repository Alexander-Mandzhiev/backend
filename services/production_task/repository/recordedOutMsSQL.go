package repository

import "context"

func (r *Repository) RecordedOutMsSQL(ctx context.Context, ids []int) error {
	// Моковая реализация удаления из MSSQL через Firebird
	if len(ids) == 0 {
		return nil
	}
	// Здесь можно добавить логику проверки ID
	return nil
}
