package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/movements"
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) InitializeProduct(ctx context.Context, req *movements.InitializeProductRequest) (*movements.InitializeProductResponse, error) {
	const op = "movements.Repository.InitializeProduct"
	logger := sl.Log.With(slog.String("op", op))
	logger.Debug("Starting batch product initialization with raw SQL")

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to start transaction: %w", op, err)
	}
	defer tx.Rollback()

	response := &movements.InitializeProductResponse{
		CreatedProducts: make([]*movements.ProductResponse, 0),
		Errors:          make([]*movements.ProductCreationError, 0),
	}

	for idx, productData := range req.GetData() {
		var productID int64
		var createdAt time.Time

		err := tx.QueryRowContext(ctx, `
            INSERT INTO products_sk 
            (part_name, nomenclature, number_frame, count_sausage_sticks, weight_sp_kg, weight_gp_kg, manufacturing_date) 
            VALUES (?, ?, ?, ?, ?, ?, ?)
            OUTPUT INSERTED.id, INSERTED.created_at`,
			productData.GetPartName(),
			productData.GetNomenclature(),
			productData.GetNumberFrame(),
			productData.CountSausageSticks,
			productData.GetWeightSpKg(),
			productData.WeightGpKg,
			productData.GetManufacturingDate().AsTime(),
		).Scan(&productID, &createdAt)

		if err != nil {
			response.Errors = append(response.Errors, &movements.ProductCreationError{
				DataIndex:    int32(idx),
				ErrorMessage: fmt.Sprintf("product creation failed: %v", err),
			})
			continue
		}

		var movementID int64
		err = tx.QueryRowContext(ctx, `
            INSERT INTO movements 
            (product_id, to_location_id, user_id, comment) 
            VALUES (?, ?, ?, ?)
            OUTPUT INSERTED.id`,
			productID,
			req.ToLocationId,
			req.UserId,
			req.Comment,
		).Scan(&movementID)

		if err != nil {
			response.Errors = append(response.Errors, &movements.ProductCreationError{
				DataIndex:    int32(idx),
				ErrorMessage: fmt.Sprintf("movement creation failed: %v", err),
			})

			if delErr := deleteProduct(tx, productID); delErr != nil {
				logger.Error("Failed to rollback product", slog.Int64("product_id", productID), sl.Err(delErr))
			}
			continue
		}

		_, err = tx.ExecContext(ctx, `
            INSERT INTO product_status_history 
            (product_id, status_id) 
            VALUES (?, ?)`,
			productID,
			req.StatusId,
		)

		if err != nil {
			response.Errors = append(response.Errors, &movements.ProductCreationError{
				DataIndex:    int32(idx),
				ErrorMessage: fmt.Sprintf("status creation failed: %v", err),
			})

			// Откат движения и продукта
			if delErr := deleteMovement(tx, movementID); delErr != nil {
				logger.Error("Failed to rollback movement", slog.Int64("movement_id", movementID), sl.Err(delErr))
			}
			if delErr := deleteProduct(tx, productID); delErr != nil {
				logger.Error("Failed to rollback product", slog.Int64("product_id", productID), sl.Err(delErr))
			}
			continue
		}

		response.CreatedProducts = append(response.CreatedProducts, &movements.ProductResponse{
			Id:                 productID,
			PartName:           productData.GetPartName(),
			Nomenclature:       productData.GetNomenclature(),
			NumberFrame:        productData.GetNumberFrame(),
			CountSausageSticks: productData.CountSausageSticks,
			WeightSpKg:         productData.GetWeightSpKg(),
			WeightGpKg:         productData.WeightGpKg,
			ManufacturingDate:  productData.GetManufacturingDate(),
			CreatedAt:          timestamppb.New(createdAt),
		})
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s: transaction commit failed: %w", op, err)
	}

	return response, nil
}

func deleteProduct(tx *sql.Tx, id int64) error {
	_, err := tx.Exec("DELETE FROM products_sk WHERE id = ?", id)
	return err
}

func deleteMovement(tx *sql.Tx, id int64) error {
	_, err := tx.Exec("DELETE FROM movements WHERE id = ?", id)
	return err
}
