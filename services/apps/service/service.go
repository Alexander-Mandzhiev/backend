package service

import (
	app "backend/protos/gen/go/apps"
	"context"
	"errors"
)

var (
	ErrAppNotFound = errors.New("app not found")
)

type AppProvider interface {
	Create(ctx context.Context, app *app.App) (int32, error)
	Apps(ctx context.Context) ([]*app.App, error)
	App(ctx context.Context, id int32) (*app.App, error)
	Update(ctx context.Context, app *app.App) error
	Delete(ctx context.Context, id int32) error
}

type Service struct {
	appProvider AppProvider
}

func New(appProvider AppProvider) *Service {
	return &Service{appProvider: appProvider}
}
