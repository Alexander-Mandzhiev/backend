package handle

import (
	"backend/protos/gen/go/sk/locations"
	"context"
)

func (s *serverAPI) ListLocation(ctx context.Context, req *locations.ListLocationsRequest) (*locations.LocationListResponse, error) {
	return s.service.Locations(ctx)
}
