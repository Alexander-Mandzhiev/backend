package apps_handle

import (
	app "backend/protos/gen/go/apps"
	"google.golang.org/grpc"
)

type AppsService struct {
	client app.AppProviderServiceClient
}

func New(conn *grpc.ClientConn) *AppsService {
	return &AppsService{client: app.NewAppProviderServiceClient(conn)}
}
