package handle

import (
	"backend/protos/gen/go/sso"
	"context"
	"google.golang.org/grpc"
)

type SSOService interface {
	SignIn(ctx context.Context, req *sso.SignInRequest) (string, error)
}

type serverAPI struct {
	sso.UnimplementedSSOServiceServer
	service SSOService
}

func Register(gRPCServer *grpc.Server, service SSOService) {
	sso.RegisterSSOServiceServer(gRPCServer, &serverAPI{service: service})
}
