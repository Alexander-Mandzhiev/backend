package handle

import (
	"backend/protos/gen/go/products_sk"
	"context"
)

func (s *serverAPI) Create(ctx context.Context, request *products_sk.CreateProductRequest) (*products_sk.ProductResponse, error) {
	return s.service.Create(ctx, request)
}
