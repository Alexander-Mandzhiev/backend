package repository

import (
	"backend/protos/gen/go/production_task"
	"context"
)

func (r *Repository) TaskInPartName(ctx context.Context, params *production_task.RequestTaskParams) ([]*production_task.Product, error) {
	// Фильтрация по PartName
	var filtered []*production_task.Product
	for _, product := range mockProducts {
		if product.PartName == params.PartName {
			filtered = append(filtered, product)
		}
	}

	return filtered, nil
}
