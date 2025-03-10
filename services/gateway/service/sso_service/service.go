package sso_service

import (
	"backend/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type SSOService struct {
	client sso.SSOServiceClient
}

func New(conn *grpc.ClientConn) *SSOService {
	return &SSOService{client: sso.NewSSOServiceClient(conn)}
}
