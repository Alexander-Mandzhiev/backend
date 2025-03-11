package sso_service

import (
	"backend/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type Service struct {
	client sso.SSOServiceClient
}

func New(conn *grpc.ClientConn) *Service {
	return &Service{client: sso.NewSSOServiceClient(conn)}
}
