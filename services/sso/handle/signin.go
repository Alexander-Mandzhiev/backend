package handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sso"
	"backend/services/sso/service"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) SignIn(ctx context.Context, req *sso.SignInRequest) (*sso.SignInResponse, error) {
	const op = "auth.SignIn"

	accessToken, err := s.service.SignIn(ctx, req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			sl.Log.Warn(op, slog.String("message", "Invalid credentials provided during sign in"))
			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		}
		sl.Log.Error(op, slog.String("message", "Failed to sign in user"), slog.Any("error", err))
		return nil, status.Error(codes.Internal, "failed to login")
	}

	return &sso.SignInResponse{AccessToken: accessToken}, nil
}
