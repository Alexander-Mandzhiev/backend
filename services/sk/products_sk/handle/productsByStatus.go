package handle

import (
	"backend/protos/gen/go/sk/products_sk"
	"context"
)

func (s *serverAPI) ProductsByStatus(ctx context.Context, request *products_sk.ProductsByStatusRequest) (*products_sk.ProductListResponse, error) {
	return s.service.ProductsByStatus(ctx, request.StatusId, request.Page, request.Count)
}
