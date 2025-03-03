package main

import (
	"backend/pkg/dbManager"
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_client"
	"backend/pkg/server/grpc_server"
	cfg "backend/services/production_task/config"
	"backend/services/production_task/handle"
	"backend/services/production_task/repository"
	"backend/services/production_task/service"
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

	srv := service.New(prv)
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
