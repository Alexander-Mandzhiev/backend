package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, request *products_sk.UpdateProductRequest) (*products_sk.ProductResponse, error) {
	op := "service.Update"
	id := request.Id
	sl.Log.Debug("Updating product", slog.Int64("id", id), slog.String("op", op))

	updatedProduct := &products_sk.ProductResponse{
		Id:                id,
		PartName:          request.PartName,
		Nomenclature:      request.Nomenclature,
		NumberFrame:       request.NumberFrame,
		WeightSpKg:        request.WeightSpKg,
		WeightGpKg:        request.WeightGpKg,
		ManufacturingDate: request.ManufacturingDate,
	}

	if err := s.productSkProvider.Update(ctx, updatedProduct); err != nil {
		sl.Log.Error("Failed to update product", sl.Err(err), slog.Int64("id", id), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Product updated successfully", slog.Int64("id", id), slog.String("op", op))
	return updatedProduct, nil
}
