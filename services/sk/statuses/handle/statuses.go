package handle

import (
	"backend/protos/gen/go/sk/statuses"
	"context"
)

func (s *serverAPI) ListStatuses(ctx context.Context, _ *statuses.ListStatusesRequest) (*statuses.StatusListResponse, error) {
	return s.service.List(ctx)
}
