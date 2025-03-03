package handle

import (
	"backend/protos/gen/go/products_sk_statuses"
	"context"
)

func (s *serverAPI) Delete(ctx context.Context, request *products_sk_statuses.DeleteProductStatusRequest) (*products_sk_statuses.DeleteProductStatusResponse, error) {
	return s.service.Delete(ctx, request)
}
