package handle

import (
	"backend/protos/gen/go/sk/location_types"
	"context"
)

func (s *serverAPI) UpdateLocationType(ctx context.Context, request *location_types.UpdateLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	return s.service.Update(ctx, request)
}
