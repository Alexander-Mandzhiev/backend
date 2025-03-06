package handle

import (
	"backend/protos/gen/go/location_types"
	"context"
)

func (s *serverAPI) Get(ctx context.Context, request *location_types.GetLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	return s.service.Get(ctx, request)
}
