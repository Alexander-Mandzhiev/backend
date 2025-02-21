package handle

import (
	"backend/protos/gen/go/statuses"
	"context"
)

func (s *serverAPI) Status(ctx context.Context, request *statuses.GetStatusRequest) (*statuses.StatusResponse, error) {
	return s.service.Status(ctx, request)
}
