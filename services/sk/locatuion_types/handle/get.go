package handle

import (
	"backend/protos/gen/go/sk/location_types"
	"context"
)

func (s *serverAPI) GetLocationType(ctx context.Context, request *location_types.GetLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	return s.service.Get(ctx, request)
}
