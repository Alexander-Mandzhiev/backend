package location_service

import (
	"backend/protos/gen/go/locations"
	"google.golang.org/grpc"
)

type LocationService struct {
	client locations.LocationServiceClient
}

func New(conn *grpc.ClientConn) *LocationService {
	return &LocationService{client: locations.NewLocationServiceClient(conn)}
}
