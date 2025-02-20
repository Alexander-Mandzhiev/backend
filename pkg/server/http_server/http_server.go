package http_server

import (
	sl "backend/pkg/logger"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"net"
)

type App struct {
	GRPCServer *grpc.Server
}

func New() *App {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived, logging.PayloadSent,
		),
	}
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			slog.Error("Recovered from panic", slog.Any("panic", p))
			return status.Errorf(codes.Internal, "internal error")
		}),
	}
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(InterceptorLogger(slog.Default()), loggingOpts...),
	))

	return &App{
		GRPCServer: gRPCServer,
	}
}

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		//l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func (a *App) MustRun(port int) {
	if err := a.Run(port); err != nil {
		sl.Log.Error("Failed to start GRPC server", slog.Any("error", err))
		panic(err)
	}
}

func (a *App) Run(port int) error {
	const op = "grpcapp.Run"
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	slog.Info("gRPC server started", slog.String("addr", listener.Addr().String()))

	if err = a.GRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *App) Shutdown() {
	slog.Info("Stopping gRPC server")
	a.GRPCServer.GracefulStop()
}
