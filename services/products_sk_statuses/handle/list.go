package handle

import (
	"backend/protos/gen/go/products_sk_statuses"
	"context"
)

func (s *serverAPI) List(ctx context.Context, request *products_sk_statuses.ListProductStatusesRequest) (*products_sk_statuses.ProductStatusListResponse, error) {
	return s.service.List(ctx, request)
}
