package service

import (
	app "backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Create(ctx context.Context, req *app.CreateRequest) (*app.CreateResponse, error) {
	op := "service.Create"
	if ctx == nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: context is nil", op)
	}
	if req.GetName() == "" || req.GetSecret() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "%s: name and secret are required", op)
	}

	id, err := s.appProvider.Create(ctx, &app.App{Name: req.GetName(), Secret: req.GetSecret()})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s: failed to create app: %v", op, err)
	}

	return &app.CreateResponse{Id: id}, nil
}
