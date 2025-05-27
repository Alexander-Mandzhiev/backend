package handle

import (
	"backend/protos/gen/go/sk/movements"
	"context"
)

func (s *serverAPI) UpdateMovement(ctx context.Context, request *movements.UpdateMovementRequest) (*movements.MovementResponse, error) {
	return s.service.Update(ctx, request)
}
