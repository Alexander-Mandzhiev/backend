package handle

import (
	"backend/protos/gen/go/sk/product_status_history"
	"context"
)

func (s *serverAPI) DeleteStatus(ctx context.Context, request *product_status_history.DeleteStatusRequest) (*product_status_history.DeleteStatusResponse, error) {
	return s.service.Delete(ctx, request)
}
