package apps_service

import (
	app "backend/protos/gen/go/apps"
	"google.golang.org/grpc"
)

type Service struct {
	client app.AppProviderServiceClient
}

func New(conn *grpc.ClientConn) *Service {
	return &Service{client: app.NewAppProviderServiceClient(conn)}
}
