package handle

import (
	"backend/protos/gen/go/sk/location_types"
	"context"
)

func (s *serverAPI) CreateLocationType(ctx context.Context, request *location_types.CreateLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	return s.service.Create(ctx, request)
}
