package handle

import (
	"backend/protos/gen/go/locations"
	"context"
)

func (s *serverAPI) List(ctx context.Context, req *locations.ListLocationsRequest) (*locations.LocationListResponse, error) {
	return s.service.Locations(ctx)
}
