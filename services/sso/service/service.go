package service

import (
	app "backend/protos/gen/go/apps"
	"backend/services/sso/models"
	"context"
	"errors"
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

// New создает новый экземпляр сервиса.
func New(userProvider UserProvider, appsClient app.AppProviderServiceClient) *Service {
	return &Service{
		userProvider: userProvider,
		appsClient:   appsClient,
	}
}
