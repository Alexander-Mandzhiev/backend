package main

import (
	"backend/pkg/dbManager"
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_server"
	cfg "backend/services/products_sk/config"
	"backend/services/products_sk/handle"
	"backend/services/products_sk/repository"
	"backend/services/products_sk/service"
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
	db, err := dbManager.OpenMSSQLConnection(cfg.Cfg.DBConfig.MSSQL.ConnectionString, cfg.Cfg.DBConfig.MSSQL.MaxOpenConnections, cfg.Cfg.DBConfig.MSSQL.MaxIdleConnections, cfg.Cfg.DBConfig.MSSQL.ConnMaxLifetime)
	if err != nil {
		log.Fatal("Failed to initialize database connection", slog.Any("error", err))
	}
	defer func() {
		if closeErr := dbManager.CloseMSSQLConnection(db); closeErr != nil {
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
