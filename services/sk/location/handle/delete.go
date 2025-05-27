package handle

import (
	"backend/protos/gen/go/sk/locations"
	"context"
)

func (s *serverAPI) DeleteLocation(ctx context.Context, req *locations.DeleteLocationRequest) (*locations.DeleteLocationResponse, error) {
	return s.service.Delete(ctx, req)
}
