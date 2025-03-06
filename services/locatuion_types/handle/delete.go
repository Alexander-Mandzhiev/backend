package handle

import (
	"backend/protos/gen/go/location_types"
	"context"
)

func (s *serverAPI) Delete(ctx context.Context, request *location_types.DeleteLocationTypeRequest) (*location_types.DeleteLocationTypeResponse, error) {
	return s.service.Delete(ctx, request)
}
