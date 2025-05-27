package handle

import (
	"backend/protos/gen/go/sk/statuses"
	"context"
)

func (s *serverAPI) DeleteStatus(ctx context.Context, request *statuses.DeleteStatusRequest) (*statuses.DeleteStatusResponse, error) {
	return s.service.Delete(ctx, request)
}
