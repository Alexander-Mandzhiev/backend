package handle

import (
	"backend/protos/gen/go/movements"
	"context"
)

func (s *serverAPI) Update(ctx context.Context, request *movements.UpdateMovementRequest) (*movements.MovementResponse, error) {
	return s.service.Update(ctx, request)
}
