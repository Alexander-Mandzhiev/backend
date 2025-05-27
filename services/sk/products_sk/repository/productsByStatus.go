package repository

import (
	"backend/protos/gen/go/sk/products_sk"
	"context"
	"database/sql"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (r *Repository) ProductsByStatus(ctx context.Context, statusID int32, page, count int32) ([]*products_sk.ProductResponse, int32, error) {
	const op = "repository.ProductsByStatus"

	offset := (page - 1) * count

	var totalItems int32
	err := r.db.QueryRowContext(ctx, `
        SELECT COUNT(*) 
        FROM products_sk p
        INNER JOIN products_sk_statuses ps ON p.id = ps.product_id
        WHERE ps.status_id = ? AND ps.active = 1 AND p.removed_at IS NULL
    `, statusID).Scan(&totalItems)
	if err != nil {
		log.Printf("[%s] Failed to count products by status: %v", op, err)
		return nil, 0, fmt.Errorf("failed to count products by status: %w", err)
	}

	if totalItems == 0 {
		log.Printf("[%s] No products found for status %d", op, statusID)
		return nil, 0, ErrProductSkNotFound
	}

	query := fmt.Sprintf(`
        SELECT 
            p.id, p.part_name, p.nomenclature, p.number_frame, 
            p.count_sausage_sticks, p.weight_sp_kg, p.weight_gp_kg, 
            p.manufacturing_date, p.created_at, p.removed_at
        FROM products_sk p
        INNER JOIN products_sk_statuses ps ON p.id = ps.product_id
        WHERE ps.status_id = ? AND ps.active = 1 AND p.removed_at IS NULL
        ORDER BY p.id
        OFFSET %d ROWS FETCH NEXT %d ROWS ONLY
    `, offset, count)

	rows, err := r.db.QueryContext(ctx, query, statusID)
	if err != nil {
		log.Printf("[%s] Failed to fetch products by status: %v", op, err)
		return nil, 0, fmt.Errorf("failed to fetch products by status: %w", err)
	}
	defer rows.Close()

	var products []*products_sk.ProductResponse

	for rows.Next() {
		var product products_sk.ProductResponse

		var (
			countSausage      sql.NullInt32
			weightGp          sql.NullFloat64
			manufacturingDate sql.NullTime
			createdAt         sql.NullTime
			removedAt         sql.NullTime
		)

		if err = rows.Scan(
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
			log.Printf("[%s] Failed to scan product row: %v", op, err)
			return nil, 0, fmt.Errorf("failed to scan product row: %w", err)
		}

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

	if err = rows.Err(); err != nil {
		log.Printf("[%s] Error occurred while iterating over rows: %v", op, err)
		return nil, 0, fmt.Errorf("error occurred while iterating over rows: %w", err)
	}

	return products, totalItems, nil
}
