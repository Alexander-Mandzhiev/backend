package service

import (
	"backend/pkg/jwt"
	sl "backend/pkg/logger"
	app "backend/protos/gen/go/apps"
	"backend/protos/gen/go/sso"
	"backend/services/sso/config"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) SignIn(ctx context.Context, req *sso.SignInRequest) (string, error) {
	const op = "service.SignIn"

	if s.userProvider == nil || s.appsClient == nil {
		return "", status.Error(codes.Internal, "service is not properly initialized")
	}

	if req.GetPassword() == 0 || req.GetAppId() == 0 {
		sl.Log.Warn(op, slog.String("message", "Missing required fields"), slog.Int("app_id", int(req.GetAppId())))
		return "", status.Error(codes.InvalidArgument, "password and app_id are required")
	}

	user, err := s.userProvider.User(ctx, int(req.GetPassword()))
	if err != nil {
		sl.Log.Error(op, slog.String("message", "Failed to get user"), slog.Any("error", err))
		return "", status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	appResp, err := s.appsClient.App(ctx, &app.GetAppRequest{Id: req.GetAppId()})
	if err != nil {
		return "", status.Errorf(codes.Internal, "failed to get app: %v", err)
	}
	if appResp == nil || appResp.App == nil {
		return "", status.Error(codes.Internal, "app response is empty")
	}

	token, err := jwt.NewToken(user.Name, user.Id, appResp.App.Secret, config.Cfg.SSO.AccessTokenTTL)
	if err != nil {
		sl.Log.Error(op, slog.String("message", "Failed to generate token"), slog.Any("error", err))
		return "", status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	return token, nil
}
