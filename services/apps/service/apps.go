package service

import (
	app "backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Apps(ctx context.Context) (*app.GetAppsResponse, error) {
	op := "service.Apps"
	if ctx == nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: context is nil", op)
	}

	apps, err := s.appProvider.Apps(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s: failed to fetch apps: %v", op, err)
	}

	return &app.GetAppsResponse{Apps: apps}, nil
}
