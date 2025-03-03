package handle

import (
	"backend/protos/gen/go/products_sk_statuses"
	"context"
)

func (s *serverAPI) Update(ctx context.Context, request *products_sk_statuses.UpdateProductStatusRequest) (*products_sk_statuses.ProductSkStatusResponse, error) {
	return s.service.Update(ctx, request)
}
