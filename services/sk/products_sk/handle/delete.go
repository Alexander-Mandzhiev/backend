package handle

import (
	"backend/protos/gen/go/sk/products_sk"
	"context"
)

func (s *serverAPI) Delete(ctx context.Context, request *products_sk.DeleteProductRequest) (*products_sk.DeleteProductResponse, error) {
	return s.service.Delete(ctx, request)
}
