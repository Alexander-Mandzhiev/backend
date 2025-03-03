package repository

import (
	"backend/protos/gen/go/production_task"
	"context"
)

func (r *Repository) Tasks(ctx context.Context, params *production_task.RequestTaskParams) ([]*production_task.Product, error) {
	// Фильтрация по параметрам
	dateStart := params.DateStart.AsTime()
	dateEnd := params.DateEnd.AsTime()

	var filtered []*production_task.Product
	for _, product := range mockProducts {
		if !product.ManufacturingDate.AsTime().Before(dateStart) && !product.ManufacturingDate.AsTime().After(dateEnd) {
			if params.Search == "" || (params.Search != "" &&
				(product.Nomenclature == params.Search || product.PartName == params.Search)) {
				filtered = append(filtered, product)
			}
		}
	}

	return filtered, nil
}
