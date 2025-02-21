package handle

import (
	"backend/protos/gen/go/locations"
	"context"
)

func (s *serverAPI) Delete(ctx context.Context, req *locations.DeleteLocationRequest) (*locations.DeleteLocationResponse, error) {
	return s.service.Delete(ctx, req)
}
