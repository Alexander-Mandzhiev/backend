package service

import (
	sl "backend/pkg/logger"
	app "backend/protos/gen/go/apps"
	"context"
	"errors"
	"log/slog"
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
	op := "service.New"
	if appProvider == nil {
		sl.Log.Error("Location provider is nil", slog.String("op", op))
		return nil
	}

	sl.Log.Info("Service initialized", slog.String("op", op))
	return &Service{appProvider: appProvider}
}
