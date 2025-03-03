package main

import (
	sl "backend/pkg/logger"
	app "backend/pkg/server/http_server"
	cfg "backend/services/gateway/config"
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

	apiServer, err := app.New()
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
