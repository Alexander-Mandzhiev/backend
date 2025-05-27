package handle

import (
	"backend/protos/gen/go/sk/location_types"
	"context"
)

func (s *serverAPI) DeleteLocationType(ctx context.Context, request *location_types.DeleteLocationTypeRequest) (*location_types.DeleteLocationTypeResponse, error) {
	return s.service.Delete(ctx, request)
}
