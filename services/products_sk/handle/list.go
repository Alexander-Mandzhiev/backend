package handle

import (
	"backend/protos/gen/go/products_sk"
	"context"
)

func (s *serverAPI) List(ctx context.Context, _ *products_sk.ListProductsRequest) (*products_sk.ProductListResponse, error) {
	return s.service.Products(ctx)
}
