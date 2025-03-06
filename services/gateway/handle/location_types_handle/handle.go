package location_types_handle

import (
	"backend/protos/gen/go/location_types"
	"google.golang.org/grpc"
)

type LocationTypesService struct {
	client location_types.LocationTypeServiceClient
}

func New(conn *grpc.ClientConn) *LocationTypesService {
	return &LocationTypesService{client: location_types.NewLocationTypeServiceClient(conn)}
}
