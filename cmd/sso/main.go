package main

import (
	"backend/pkg/dbManager"
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_client"
	"backend/pkg/server/grpc_server"
	app_provider "backend/protos/gen/go/apps"
	cfg "backend/services/sso/config"
	"backend/services/sso/handle"
	"backend/services/sso/repository"
	"backend/services/sso/service"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg.Initialize()
	sl.SetupLogger(cfg.Cfg.Env)
	sl.Log.Info("Starting service sso", slog.String("address", cfg.Cfg.GRPCServer.Address), slog.Int("port", cfg.Cfg.GRPCServer.Port))
	sl.Log.Debug("Debug messages are enabled")

	clientManager := grpc_client.NewGRPCClientManager()
	defer clientManager.CloseAll()

	db, err := dbManager.OpenFirebirdConnection(cfg.Cfg.DBConfig.Firebird.ConnectionString, cfg.Cfg.DBConfig.Firebird.MaxOpenConnections, cfg.Cfg.DBConfig.Firebird.MaxIdleConnections, cfg.Cfg.DBConfig.Firebird.ConnMaxLifetime)
	if err != nil {
		log.Fatal("Failed to initialize database connection", slog.Any("error", err))
	}
	defer func() {
		if closeErr := dbManager.CloseFirebirdConnection(db); closeErr != nil {
			log.Fatalf("Failed to close database connection: %v", closeErr)
		}
	}()

	prv, err := repository.New(db)
	if err != nil {
		log.Fatal("Failed to initialize repository", slog.Any("error", err))
	}

	appsAddr := fmt.Sprintf("%s:%d", cfg.Cfg.Services.AppsService.Address, cfg.Cfg.Services.AppsService.Port)
	appsConn, err := clientManager.GetClientConn(appsAddr)
	if err != nil {
		log.Fatalf("Failed to connect to apps service: %v", err)
	}

	srv := service.New(prv, app_provider.NewAppProviderServiceClient(appsConn))
	application := grpc_server.New()
	handle.Register(application.GRPCServer, srv)

	go func() {
		application.MustRun(cfg.Cfg.GRPCServer.Port)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Shutdown()
	sl.Log.Info("Gracefully stopped")
}
