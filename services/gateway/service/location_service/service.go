package location_service

import (
	"backend/protos/gen/go/locations"
	"google.golang.org/grpc"
)

type Service struct {
	client locations.LocationServiceClient
}

func New(conn *grpc.ClientConn) *Service {
	return &Service{client: locations.NewLocationServiceClient(conn)}
}
