package handle

import (
	"backend/protos/gen/go/sk/movements"
	"context"
)

func (s *serverAPI) DeleteMovement(ctx context.Context, request *movements.DeleteMovementRequest) (*movements.DeleteMovementResponse, error) {
	return s.service.Delete(ctx, request)
}
