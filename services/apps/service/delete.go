package service

import (
	app_provider "backend/protos/gen/go/apps"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Delete(ctx context.Context, req *app_provider.DeleteRequest) (*app_provider.DeleteResponse, error) {
	op := "service.Delete"
	if ctx == nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: context is nil", op)
	}
	if req.GetId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%s: invalid app ID", op)
	}

	err := s.appProvider.Delete(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, ErrAppNotFound) {
			return nil, status.Errorf(codes.NotFound, "%s: app not found", op)
		}
		return nil, status.Errorf(codes.Internal, "%s: failed to delete app: %v", op, err)
	}

	return &app_provider.DeleteResponse{Success: true}, nil
}
