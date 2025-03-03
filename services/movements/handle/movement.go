package handle

import (
	"backend/protos/gen/go/movements"
	"context"
)

func (s *serverAPI) Movement(ctx context.Context, request *movements.GetMovementRequest) (*movements.MovementResponse, error) {
	return s.service.Movement(ctx, request)
}
