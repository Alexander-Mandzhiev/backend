package handle

import (
	"backend/protos/gen/go/locations"
	"context"
)

func (s *serverAPI) Location(ctx context.Context, req *locations.GetLocationRequest) (*locations.LocationResponse, error) {
	return s.service.Location(ctx, req)
}
