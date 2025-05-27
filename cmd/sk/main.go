package main

import (
	"backend/pkg/dbManager"
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_client"
	"backend/pkg/server/grpc_server"
	cfg "backend/services/sk/config"
	"backend/services/sk/location/handle"
	"backend/services/sk/location/repository"
	"backend/services/sk/location/service"

	handleLT "backend/services/sk/locatuion_types/handle"
	repositoryLT "backend/services/sk/locatuion_types/repository"
	serviceLT "backend/services/sk/locatuion_types/service"

	handleS "backend/services/sk/statuses/handle"
	repositoryS "backend/services/sk/statuses/repository"
	serviceS "backend/services/sk/statuses/service"
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

	db, err := dbManager.OpenMSSQLConnection(cfg.Cfg.DBConfig.MSSQL.ConnectionString, cfg.Cfg.DBConfig.MSSQL.MaxOpenConnections, cfg.Cfg.DBConfig.MSSQL.MaxIdleConnections, cfg.Cfg.DBConfig.MSSQL.ConnMaxLifetime)
	if err != nil {
		log.Fatal("Failed to initialize database connection", slog.Any("error", err))
	}
	defer func() {
		if closeErr := dbManager.CloseMSSQLConnection(db); closeErr != nil {
			log.Fatalf("Failed to close database connection: %v", closeErr)
		}
	}()

	prvLT, err := repositoryLT.New(db)
	if err != nil {
		log.Fatal("Failed to initialize repository", slog.Any("error", err))
	}
	srvLT := serviceLT.New(prvLT)

	prvL, err := repository.New(db)
	if err != nil {
		log.Fatal("Failed to initialize repository", slog.Any("error", err))
	}
	srvL := service.New(prvL, prvLT)

	prvS, err := repositoryS.New(db)
	if err != nil {
		log.Fatal("Failed to initialize repository", slog.Any("error", err))
	}
	srvS := serviceS.New(prvS)

	application := grpc_server.New()
	handleLT.Register(application.GRPCServer, srvLT)
	handle.Register(application.GRPCServer, srvL)
	handleS.Register(application.GRPCServer, srvS)
	go func() {
		application.MustRun(cfg.Cfg.GRPCServer.Port)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Shutdown()
	sl.Log.Info("Gracefully stopped")
}
