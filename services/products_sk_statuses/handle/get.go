package handle

import (
	"backend/protos/gen/go/products_sk_statuses"
	"context"
)

func (s *serverAPI) ProductSkStatus(ctx context.Context, request *products_sk_statuses.GetProductStatusRequest) (*products_sk_statuses.ProductSkStatusResponse, error) {
	return s.service.Get(ctx, request)
}
