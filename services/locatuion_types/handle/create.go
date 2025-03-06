package handle

import (
	"backend/protos/gen/go/location_types"
	"context"
)

func (s *serverAPI) Create(ctx context.Context, request *location_types.CreateLocationTypeRequest) (*location_types.LocationTypeResponse, error) {
	return s.service.Create(ctx, request)
}
