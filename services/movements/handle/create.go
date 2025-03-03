package handle

import (
	"backend/protos/gen/go/movements"
	"context"
)

func (s *serverAPI) Create(ctx context.Context, request *movements.CreateMovementRequest) (*movements.MovementResponse, error) {
	return s.service.Create(ctx, request)
}
