package http_app

import (
	cfg "backend/services/gateway/config"
	"backend/services/gateway/handle"
	"context"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
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
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*", "http://*"},
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
	return s.httpserver.Shutdown(ctx)
}
