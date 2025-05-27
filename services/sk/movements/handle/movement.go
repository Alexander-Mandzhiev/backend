package handle

import (
	"backend/protos/gen/go/sk/movements"
	"context"
)

func (s *serverAPI) GetMovement(ctx context.Context, request *movements.GetMovementRequest) (*movements.MovementResponse, error) {
	return s.service.Movement(ctx, request)
}
