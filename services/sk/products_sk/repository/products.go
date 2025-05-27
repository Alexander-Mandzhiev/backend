package repository

import (
	"backend/protos/gen/go/sk/products_sk"
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (r *Repository) Products(ctx context.Context, page, count int32) ([]*products_sk.ProductResponse, int32, error) {
	offset := (page - 1) * count

	var totalItems int32
	if err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM products_sk").Scan(&totalItems); err != nil {
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}

	query := fmt.Sprintf(`
        SELECT 
            id, part_name, nomenclature, number_frame, count_sausage_sticks,
            weight_sp_kg, weight_gp_kg, manufacturing_date, created_at, removed_at
        FROM products_sk
        ORDER BY id
        OFFSET %d ROWS FETCH NEXT %d ROWS ONLY`, offset, count)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch products: %w", err)
	}
	defer rows.Close()

	var products []*products_sk.ProductResponse
	for rows.Next() {
		var product products_sk.ProductResponse
		var (
			manufacturingDate sql.NullTime
			createdAt         sql.NullTime
			removedAt         sql.NullTime
			countSausage      sql.NullInt32
			weightGp          sql.NullFloat64
		)

		if err := rows.Scan(
			&product.Id,
			&product.PartName,
			&product.Nomenclature,
			&product.NumberFrame,
			&countSausage,
			&product.WeightSpKg,
			&weightGp,
			&manufacturingDate,
			&createdAt,
			&removedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan product row: %w", err)
		}

		// Обработка nullable полей
		if countSausage.Valid {
			product.CountSausageSticks = &countSausage.Int32
		}
		if weightGp.Valid {
			product.WeightGpKg = &weightGp.Float64
		}
		if manufacturingDate.Valid {
			product.ManufacturingDate = timestamppb.New(manufacturingDate.Time)
		}
		if createdAt.Valid {
			product.CreatedAt = timestamppb.New(createdAt.Time)
		}
		if removedAt.Valid {
			product.RemovedAt = timestamppb.New(removedAt.Time)
		}

		products = append(products, &product)
	}

	return products, totalItems, nil
}
