package handle

import (
	"backend/protos/gen/go/statuses"
	"context"
)

func (s *serverAPI) Create(ctx context.Context, request *statuses.CreateStatusRequest) (*statuses.StatusResponse, error) {
	return s.service.Create(ctx, request)
}
