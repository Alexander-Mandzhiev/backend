package handle

import (
	"backend/protos/gen/go/sk/product_status_history"
	"context"
)

func (s *serverAPI) GetStatus(ctx context.Context, request *product_status_history.GetStatusRequest) (*product_status_history.ProductSkStatusResponse, error) {
	return s.service.Get(ctx, request)
}
