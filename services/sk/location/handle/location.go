package handle

import (
	"backend/protos/gen/go/sk/locations"
	"context"
)

func (s *serverAPI) GetLocation(ctx context.Context, req *locations.GetLocationRequest) (*locations.LocationResponse, error) {
	return s.service.Location(ctx, req)
}
