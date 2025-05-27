package handle

import (
	"backend/protos/gen/go/sk/location_types"
	"context"
	"google.golang.org/grpc"
)

type LocationTypesService interface {
	Create(ctx context.Context, request *location_types.CreateLocationTypeRequest) (*location_types.LocationTypeResponse, error)
	Get(ctx context.Context, request *location_types.GetLocationTypeRequest) (*location_types.LocationTypeResponse, error)
	Update(ctx context.Context, request *location_types.UpdateLocationTypeRequest) (*location_types.LocationTypeResponse, error)
	Delete(ctx context.Context, request *location_types.DeleteLocationTypeRequest) (*location_types.DeleteLocationTypeResponse, error)
	List(ctx context.Context, _ *location_types.ListLocationTypesRequest) (*location_types.LocationTypeListResponse, error)
}

type serverAPI struct {
	location_types.UnimplementedLocationTypeServiceServer
	service LocationTypesService
}

func Register(gRPCServer *grpc.Server, service LocationTypesService) {
	location_types.RegisterLocationTypeServiceServer(gRPCServer, &serverAPI{service: service})
}
