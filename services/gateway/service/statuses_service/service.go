package statuses_service

import (
	"backend/protos/gen/go/statuses"
	"google.golang.org/grpc"
)

type Service struct {
	client statuses.StatusServiceClient
}

func New(conn *grpc.ClientConn) *Service {
	return &Service{client: statuses.NewStatusServiceClient(conn)}
}
