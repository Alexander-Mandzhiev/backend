package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/products_sk"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, request *products_sk.CreateProductRequest) (*products_sk.ProductResponse, error) {
	op := "service.Create"
	sl.Log.Debug("Creating new product", slog.String("op", op), slog.Int64("id", request.Id))

	newProduct := &products_sk.ProductResponse{
		Id:                request.Id,
		PartName:          request.PartName,
		Nomenclature:      request.Nomenclature,
		NumberFrame:       request.NumberFrame,
		WeightSpKg:        request.WeightSpKg,
		WeightGpKg:        request.WeightGpKg,
		ManufacturingDate: request.ManufacturingDate,
	}

	if newProduct.ManufacturingDate == nil {
		newProduct.ManufacturingDate = timestamppb.Now()
	}

	err := s.productSkProvider.Create(ctx, newProduct)
	if err != nil {
		sl.Log.Error("Failed to create product", sl.Err(err), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Product created successfully", slog.Int64("id", newProduct.Id), slog.String("op", op))
	return newProduct, nil
}
