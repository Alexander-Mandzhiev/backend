package handle

import (
	"backend/protos/gen/go/sk/products_sk"
	"context"
)

func (s *serverAPI) Product(ctx context.Context, request *products_sk.GetProductRequest) (*products_sk.ProductResponse, error) {
	return s.service.Product(ctx, request)
}
