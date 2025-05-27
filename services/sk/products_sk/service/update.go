package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sk/products_sk"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, request *products_sk.UpdateProductRequest) (*products_sk.ProductResponse, error) {
	const op = "service.Update"
	logger := sl.Log.With(slog.String("op", op))

	currentProduct, err := s.productSkProvider.Product(ctx, request.Id)
	if err != nil {
		logger.Error("product not found", slog.Int64("id", request.Id), sl.Err(err))
		return nil, fmt.Errorf("product %d not found", request.Id)
	}

	updatedProduct := &products_sk.ProductResponse{
		Id:                request.Id,
		PartName:          currentProduct.PartName,
		Nomenclature:      currentProduct.Nomenclature,
		NumberFrame:       currentProduct.NumberFrame,
		WeightSpKg:        currentProduct.WeightSpKg,
		WeightGpKg:        currentProduct.WeightGpKg,
		ManufacturingDate: currentProduct.ManufacturingDate,
	}

	if request.PartName != "" {
		updatedProduct.PartName = request.PartName
	}
	if request.Nomenclature != "" {
		updatedProduct.Nomenclature = request.Nomenclature
	}
	if request.NumberFrame != 0 {
		updatedProduct.NumberFrame = request.NumberFrame
	}
	if request.CountSausageSticks != 0 {
		updatedProduct.CountSausageSticks = &request.CountSausageSticks
	}
	if request.WeightSpKg != 0 {
		updatedProduct.WeightSpKg = request.WeightSpKg
	}
	if request.WeightGpKg != 0 {
		updatedProduct.WeightGpKg = &request.WeightGpKg
	}
	if request.ManufacturingDate != nil {
		updatedProduct.ManufacturingDate = request.ManufacturingDate
	}

	if err = s.productSkProvider.Update(ctx, updatedProduct); err != nil {
		logger.Error("update failed", slog.Int64("id", request.Id), sl.Err(err))
		return nil, fmt.Errorf("update failed: %w", err)
	}

	logger.Info("product updated", slog.Int64("id", request.Id))
	return updatedProduct, nil
}
