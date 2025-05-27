package handle

import (
	"backend/protos/gen/go/sk/product_status_history"
	"context"
)

func (s *serverAPI) CreateStatus(ctx context.Context, request *product_status_history.CreateStatusRequest) (*product_status_history.ProductSkStatusResponse, error) {
	return s.service.Create(ctx, request)
}
