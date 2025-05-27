package handle

import (
	"backend/protos/gen/go/sk/statuses"
	"context"
)

func (s *serverAPI) CreateStatus(ctx context.Context, request *statuses.CreateStatusRequest) (*statuses.StatusResponse, error) {
	return s.service.Create(ctx, request)
}
