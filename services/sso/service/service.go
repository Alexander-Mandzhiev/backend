package service

import (
	sl "backend/pkg/logger"
	app "backend/protos/gen/go/apps"
	"backend/services/sso/models"
	"context"
	"errors"
	"log/slog"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Service struct {
	userProvider UserProvider
	appsClient   app.AppProviderServiceClient
}

type UserProvider interface {
	User(ctx context.Context, password int) (models.User, error)
}

func New(userProvider UserProvider, appsClient app.AppProviderServiceClient) *Service {
	op := "service.New"
	if userProvider == nil {
		sl.Log.Error("User provider is nil", slog.String("op", op))
		return nil
	}

	if appsClient == nil {
		sl.Log.Error("App provider is nil", slog.String("op", op))
		return nil
	}
	sl.Log.Info("Service initialized", slog.String("op", op))
	return &Service{
		userProvider: userProvider,
		appsClient:   appsClient,
	}
}
