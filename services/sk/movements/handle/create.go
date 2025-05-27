package handle

import (
	"backend/protos/gen/go/sk/movements"
	"context"
)

func (s *serverAPI) CreateMovement(ctx context.Context, request *movements.CreateMovementRequest) (*movements.MovementResponse, error) {
	return s.service.Create(ctx, request)
}
