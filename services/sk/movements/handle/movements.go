package handle

import (
	"backend/protos/gen/go/sk/movements"
	"context"
)

func (s *serverAPI) ListMovements(ctx context.Context, request *movements.ListMovementsRequest) (*movements.MovementListResponse, error) {
	return s.service.Movements(ctx, request)
}
