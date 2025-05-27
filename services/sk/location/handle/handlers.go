package handle

import (
	"backend/protos/gen/go/sk/locations"
	"context"
	"google.golang.org/grpc"
)

type LocationService interface {
	Create(ctx context.Context, request *locations.CreateLocationRequest) (*locations.LocationResponse, error)
	Location(ctx context.Context, request *locations.GetLocationRequest) (*locations.LocationResponse, error)
	Update(ctx context.Context, request *locations.UpdateLocationRequest) (*locations.LocationResponse, error)
	Delete(ctx context.Context, request *locations.DeleteLocationRequest) (*locations.DeleteLocationResponse, error)
	Locations(ctx context.Context) (*locations.LocationListResponse, error)
}

type serverAPI struct {
	locations.UnimplementedLocationServiceServer
	service LocationService
}

func Register(gRPCServer *grpc.Server, service LocationService) {
	locations.RegisterLocationServiceServer(gRPCServer, &serverAPI{service: service})
}
