package handle

import (
	"backend/protos/gen/go/sk/locations"
	"context"
)

func (s *serverAPI) UpdateLocation(ctx context.Context, req *locations.UpdateLocationRequest) (*locations.LocationResponse, error) {
	return s.service.Update(ctx, req)
}
