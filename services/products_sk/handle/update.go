package handle

import (
	"backend/protos/gen/go/products_sk"
	"context"
)

func (s *serverAPI) Update(ctx context.Context, request *products_sk.UpdateProductRequest) (*products_sk.ProductResponse, error) {
	return s.service.Update(ctx, request)
}
