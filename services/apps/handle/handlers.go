package handle

import (
	app "backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc"
)

type AppsService interface {
	Create(ctx context.Context, in *app.CreateRequest) (*app.CreateResponse, error)
	Apps(ctx context.Context) (*app.GetAppsResponse, error)
	App(ctx context.Context, in *app.GetAppRequest) (*app.GetAppResponse, error)
	Update(ctx context.Context, in *app.UpdateRequest) (*app.UpdateResponse, error)
	Delete(ctx context.Context, in *app.DeleteRequest) (*app.DeleteResponse, error)
}

type serverAPI struct {
	app.UnimplementedAppProviderServiceServer
	service AppsService
}

func Register(gRPCServer *grpc.Server, service AppsService) {
	app.RegisterAppProviderServiceServer(gRPCServer, &serverAPI{service: service})
}
