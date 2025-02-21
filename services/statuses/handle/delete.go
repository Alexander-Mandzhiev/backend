package handle

import (
	"backend/protos/gen/go/statuses"
	"context"
)

func (s *serverAPI) Delete(ctx context.Context, request *statuses.DeleteStatusRequest) (*statuses.DeleteStatusResponse, error) {
	return s.service.Delete(ctx, request)
}
