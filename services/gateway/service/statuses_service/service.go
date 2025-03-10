package statuses_service

import (
	"backend/protos/gen/go/statuses"
	"google.golang.org/grpc"
)

type StatusesService struct {
	client statuses.StatusServiceClient
}

func New(conn *grpc.ClientConn) *StatusesService {
	return &StatusesService{client: statuses.NewStatusServiceClient(conn)}
}
