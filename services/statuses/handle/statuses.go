package handle

import (
	"backend/protos/gen/go/statuses"
	"context"
)

func (s *serverAPI) List(ctx context.Context, _ *statuses.ListStatusesRequest) (*statuses.StatusListResponse, error) {
	return s.service.List(ctx)
}
