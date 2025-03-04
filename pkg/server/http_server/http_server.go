package http_app

import (
	cfg "backend/services/gateway/config"
	"backend/services/gateway/handle"
	"context"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

type APIServer struct {
	httpserver *http.Server
	gateway    *handle.Handler
}

func New() (*APIServer, error) {
	handler, err := initializeHandler()
	if err != nil {
		log.Fatalf("Failed to initialize handler: %v", err)
	}

	return &APIServer{gateway: handler}, nil
}

func (s *APIServer) Start() error {
	frontendAddr := cfg.Cfg.Frontend.Addr
	if frontendAddr == "" {
		log.Fatalf("Frontend address is not configured")
	}

	allowedOrigins := []string{
		fmt.Sprintf("http://%s", frontendAddr),
		fmt.Sprintf("https://%s", frontendAddr),
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With", "X-API-Key", "X-Csrf-Token"},
	})
	handlerWithCORS := c.Handler(s.gateway.InitRouters())

	s.httpserver = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", cfg.Cfg.Address, cfg.Cfg.Port),
		Handler:        handlerWithCORS,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    cfg.Cfg.Timeout,
		WriteTimeout:   cfg.Cfg.Timeout,
		IdleTimeout:    cfg.Cfg.IdleTimeout,
	}

	return s.httpserver.ListenAndServe()
}

func (s *APIServer) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	log.Println("Shutting down HTTP server...")
	err := s.httpserver.Shutdown(ctx)
	if err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}
	return err
}
