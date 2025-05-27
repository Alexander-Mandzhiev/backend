package handle

import (
	"backend/protos/gen/go/sk/statuses"
	"context"
)

func (s *serverAPI) GetStatus(ctx context.Context, request *statuses.GetStatusRequest) (*statuses.StatusResponse, error) {
	return s.service.Status(ctx, request)
}
