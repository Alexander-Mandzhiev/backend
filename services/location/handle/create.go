package handle

import (
	"backend/protos/gen/go/locations"
	"context"
)

func (s *serverAPI) Create(ctx context.Context, req *locations.CreateLocationRequest) (*locations.LocationResponse, error) {
	return s.service.Create(ctx, req)
}
