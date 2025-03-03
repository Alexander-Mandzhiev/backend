package handle

import (
	"backend/protos/gen/go/movements"
	"context"
)

func (s *serverAPI) List(ctx context.Context, request *movements.ListMovementsRequest) (*movements.MovementListResponse, error) {
	return s.service.Movements(ctx, request)
}
