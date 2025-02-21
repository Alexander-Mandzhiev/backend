package handle

import (
	"backend/protos/gen/go/statuses"
	"context"
)

func (s *serverAPI) Update(ctx context.Context, request *statuses.UpdateStatusRequest) (*statuses.StatusResponse, error) {
	return s.service.Update(ctx, request)
}
