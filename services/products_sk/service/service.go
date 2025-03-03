package service

import (
	"backend/protos/gen/go/products_sk"
	"context"
	"errors"
)

var ErrProductSkNotFound = errors.New("product not found")

type ProductSkProvider interface {
	Create(ctx context.Context, product *products_sk.ProductResponse) error
	Product(ctx context.Context, id int64) (*products_sk.ProductResponse, error)
	Update(ctx context.Context, product *products_sk.ProductResponse) error
	Delete(ctx context.Context, id int64) error
	Products(ctx context.Context) ([]*products_sk.ProductResponse, error)
	ProductsByStatus(ctx context.Context, statusID int32) ([]*products_sk.ProductResponse, error)
}

type Service struct {
	productSkProvider ProductSkProvider
}

func New(productSkProvider ProductSkProvider) *Service {
	return &Service{productSkProvider: productSkProvider}
}
