package handle

import (
	"backend/protos/gen/go/movements"
	"context"
)

func (s *serverAPI) Delete(ctx context.Context, request *movements.DeleteMovementRequest) (*movements.DeleteMovementResponse, error) {
	return s.service.Delete(ctx, request)
}
