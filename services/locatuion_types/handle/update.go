package handle

import (
	"backend/protos/gen/go/location_types"
	"context"
)

func (s *serverAPI) Update(ctx context.Context, request *location_types.UpdateLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	return s.service.Update(ctx, request)
}
