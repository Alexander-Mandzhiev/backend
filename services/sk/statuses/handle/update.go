package handle

import (
	"backend/protos/gen/go/sk/statuses"
	"context"
)

func (s *serverAPI) UpdateStatus(ctx context.Context, request *statuses.UpdateStatusRequest) (*statuses.StatusResponse, error) {
	return s.service.Update(ctx, request)
}
