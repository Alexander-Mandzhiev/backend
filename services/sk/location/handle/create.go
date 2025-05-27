package handle

import (
	"backend/protos/gen/go/sk/locations"
	"context"
)

func (s *serverAPI) CreateLocation(ctx context.Context, req *locations.CreateLocationRequest) (*locations.LocationResponse, error) {
	return s.service.Create(ctx, req)
}
