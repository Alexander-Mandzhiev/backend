package main

import (
	sl "backend/pkg/logger"
	app "backend/pkg/server/http_server"
	cfg "backend/services/gateway/config"
	"backend/services/gateway/handle"
	"backend/services/gateway/handle/apps"
	"backend/services/gateway/handle/location"
	"backend/services/gateway/handle/location_types"
	"backend/services/gateway/handle/production_task"
	"backend/services/gateway/handle/sso"
	"backend/services/gateway/handle/statuses"
	"backend/services/gateway/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg.Initialize()
	sl.SetupLogger(cfg.Cfg.Env)
	log.Printf("Starting API Gateway on %s:%d", cfg.Cfg.Address, cfg.Cfg.Port)

	connSet, err := service.CreateGRPCClients()
	if err != nil {
		log.Fatalf("Failed to initialize gRPC clients: %v", err)
	}
	defer connSet.CloseAll()

	gatewayService := service.New(connSet.SsoConn, connSet.AppsConn, connSet.LocationsConn, connSet.LocationTypesConn,
		connSet.MovementsConn, connSet.ProductionTasksConn, connSet.ProductSKConn, connSet.ProductsSKStatusesConn, connSet.StatusesConn)

	ssoHandler := sso_handle.New(gatewayService.SSOClient)
	appsHandler := apps_handle.New(gatewayService.AppsClient)
	productionTasksHandle := production_task_handle.New(gatewayService.ProductionTasksClient)

	locationTypesHandle := location_types_handle.New(gatewayService.LocationTypesClient)
	locationsHandler := location_handle.New(gatewayService.LocationsClient)
	statusesHandler := statuses_handle.New(gatewayService.StatusesClient)

	serverAPI := handle.New(appsHandler, ssoHandler, statusesHandler, locationsHandler, locationTypesHandle, productionTasksHandle)

	apiServer, err := app.New(serverAPI)
	if err != nil {
		log.Fatalf("Failed to initialize API Gateway: %v", err)
	}

	go func() {
		if err = apiServer.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("API Gateway failed: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = apiServer.Shutdown(ctx); err != nil {
		log.Fatalf("API Gateway shutdown failed: %v", err)
	}

	log.Println("API Gateway gracefully stopped")
}
