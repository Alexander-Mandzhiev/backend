package handle

import (
	"backend/protos/gen/go/sk/movements"
	"context"
)

func (s *serverAPI) InitializeProduct(ctx context.Context, request *movements.InitializeProductRequest) (*movements.InitializeProductResponse, error) {
	return s.service.InitializeProduct(ctx, request)
}
