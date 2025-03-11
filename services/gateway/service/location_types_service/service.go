package location_types_service

import (
	"backend/protos/gen/go/location_types"
	"google.golang.org/grpc"
)

type Service struct {
	client location_types.LocationTypeServiceClient
}

func New(conn *grpc.ClientConn) *Service {
	return &Service{client: location_types.NewLocationTypeServiceClient(conn)}
}
