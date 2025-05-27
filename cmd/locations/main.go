package main

import (
	"backend/pkg/dbManager"
	sl "backend/pkg/logger"
	"backend/pkg/server/grpc_client"
	"backend/pkg/server/grpc_server"
	cfg "backend/services/sk/config"

	handleL "backend/services/sk/location/handle"
	repositoryL "backend/services/sk/location/repository"
	serviceL "backend/services/sk/location/service"

	handleLT "backend/services/sk/locatuion_types/handle"
	repositoryLT "backend/services/sk/locatuion_types/repository"
	serviceLT "backend/services/sk/locatuion_types/service"

	handleS "backend/services/sk/statuses/handle"
	repositoryS "backend/services/sk/statuses/repository"
	serviceS "backend/services/sk/statuses/service"

	handleSK "backend/services/sk/products_sk/handle"
	repositorySK "backend/services/sk/products_sk/repository"
	serviceSK "backend/services/sk/products_sk/service"

	handleM "backend/services/sk/movements/handle"
	repositoryM "backend/services/sk/movements/repository"
	serviceM "backend/services/sk/movements/service"

	handleSKH "backend/services/sk/product_status_history/handle"
	repositorySKH "backend/services/sk/product_status_history/repository"
	serviceSKH "backend/services/sk/product_status_history/service"
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

	grpcManager := grpc_client.NewGRPCClientManager(sl.Log)
	defer func() {
		if err := grpcManager.CloseAll(); err != nil {
			sl.Log.Error("Failed to close GRPC connections", sl.Err(err))
		}
	}()

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

	prvL, err := repositoryL.New(db)
	if err != nil {
		log.Fatal("Failed to initialize repository", slog.Any("error", err))
	}

	prvS, err := repositoryS.New(db)
	if err != nil {
		log.Fatal("Failed to initialize repository", slog.Any("error", err))
	}

	prvSK, err := repositorySK.New(db)
	if err != nil {
		log.Fatal("Failed to initialize repository", slog.Any("error", err))
	}

	prvM, err := repositoryM.New(db)
	if err != nil {
		log.Fatal("Failed to initialize repository", slog.Any("error", err))
	}
	prvSKH, err := repositorySKH.New(db)
	if err != nil {
		log.Fatal("Failed to initialize repository", slog.Any("error", err))
	}

	srvSKH := serviceSKH.New(prvSKH)
	srvS := serviceS.New(prvS)
	srvL := serviceL.New(prvL, prvLT)
	srvLT := serviceLT.New(prvLT)
	srvM := serviceM.New(prvM, prvSK, prvS, prvSKH, prvL)
	srvSK := serviceSK.New(prvSK, prvM, prvS, prvSKH, prvL)

	application := grpc_server.New()

	handleLT.Register(application.GRPCServer, srvLT)
	handleL.Register(application.GRPCServer, srvL)
	handleS.Register(application.GRPCServer, srvS)
	handleM.Register(application.GRPCServer, srvM)
	handleSK.Register(application.GRPCServer, srvSK)
	handleSKH.Register(application.GRPCServer, srvSKH)

	go func() {
		application.MustRun(cfg.Cfg.GRPCServer.Port)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Shutdown()
	sl.Log.Info("Gracefully stopped")
}
