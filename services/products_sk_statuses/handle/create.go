package handle

import (
	"backend/protos/gen/go/products_sk_statuses"
	"context"
)

func (s *serverAPI) Create(ctx context.Context, request *products_sk_statuses.CreateProductStatusRequest) (*products_sk_statuses.ProductSkStatusResponse, error) {
	return s.service.Create(ctx, request)
}
