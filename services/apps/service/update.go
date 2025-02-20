package service

import (
	app "backend/protos/gen/go/apps"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Update(ctx context.Context, req *app.UpdateRequest) (*app.UpdateResponse, error) {
	op := "service.Update"
	if ctx == nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: context is nil", op)
	}
	if req.GetId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%s: invalid app ID", op)
	}
	if req.GetName() == "" || req.GetSecret() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "%s: name and secret are required", op)
	}

	err := s.appProvider.Update(ctx, &app.App{Id: req.GetId(), Name: req.GetName(), Secret: req.GetSecret()})
	if err != nil {
		if errors.Is(err, ErrAppNotFound) {
			return nil, status.Errorf(codes.NotFound, "%s: app not found", op)
		}
		return nil, status.Errorf(codes.Internal, "%s: failed to update app: %v", op, err)
	}

	return &app.UpdateResponse{Success: true}, nil
}
