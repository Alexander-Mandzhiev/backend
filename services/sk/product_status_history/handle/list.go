package handle

import (
	"backend/protos/gen/go/sk/product_status_history"
	"context"
)

func (s *serverAPI) ListStatusesByProduct(ctx context.Context, request *product_status_history.ListStatusesByProductRequest) (*product_status_history.StatusListResponse, error) {
	return s.service.List(ctx, request)
}
