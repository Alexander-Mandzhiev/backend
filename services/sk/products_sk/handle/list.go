package handle

import (
	"backend/protos/gen/go/sk/products_sk"
	"context"
)

func (s *serverAPI) List(ctx context.Context, req *products_sk.ListProductsRequest) (*products_sk.ProductListResponse, error) {
	return s.service.Products(ctx, req.GetPage(), req.GetCount())
}
