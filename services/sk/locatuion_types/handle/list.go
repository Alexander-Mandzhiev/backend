package handle

import (
	"backend/protos/gen/go/sk/location_types"
	"context"
)

func (s *serverAPI) ListLocationType(ctx context.Context, request *location_types.ListLocationTypesRequest) (*location_types.LocationTypeListResponse, error) {
	return s.service.List(ctx, request)
}
