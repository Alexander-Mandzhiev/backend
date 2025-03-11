package sso_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/sso"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) SignIn(ctx context.Context, req *sso.SignInRequest) (string, error) {
	op := "sso.SignIn"
	if req.Password == 0 || req.AppId <= 0 {
		sl.Log.Warn("Missing required fields", slog.Int64("password", req.Password), slog.Int("app_id", int(req.AppId)), slog.String("op", op))
		return "", errors.New("missing required fields: password and app_id")
	}

	sl.Log.Info("Processing sign in request", slog.Int64("password", req.Password), slog.Int("app_id", int(req.AppId)), slog.String("op", op))

	resp, err := s.client.SignIn(ctx, req)
	if err != nil {
		sl.Log.Error("Sign in failed", sl.Err(err), slog.Int("app_id", int(req.AppId)), slog.String("op", op))
		return "", err
	}

	return resp.GetAccessToken(), nil
}
