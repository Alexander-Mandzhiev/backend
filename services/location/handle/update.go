package handle

import (
	"backend/protos/gen/go/locations"
	"context"
)

func (s *serverAPI) Update(ctx context.Context, req *locations.UpdateLocationRequest) (*locations.LocationResponse, error) {
	return s.service.Update(ctx, req)
}
