package http_app

import (
	"backend/pkg/server/grpc_client"
	cfg "backend/services/gateway/config"

	"backend/services/gateway/handle"
	"fmt"
)

func initializeHandler() (*handle.Handler, error) {
	clientManager := grpc_client.NewGRPCClientManager()
	ssoConn, err := clientManager.GetClientConn(cfg.Cfg.Services.Sso)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SSO service: %w", err)
	}
	appsConn, err := clientManager.GetClientConn(cfg.Cfg.Services.Apps)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Apps service: %w", err)
	}
	locationsConn, err := clientManager.GetClientConn(cfg.Cfg.Services.Locations)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Locations service: %w", err)
	}
	movementsConn, err := clientManager.GetClientConn(cfg.Cfg.Services.Movements)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Movements service: %w", err)
	}
	productionTasksConn, err := clientManager.GetClientConn(cfg.Cfg.Services.ProductionTasks)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Production Tasks service: %w", err)
	}
	productSKConn, err := clientManager.GetClientConn(cfg.Cfg.Services.ProductSK)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Product SK service: %w", err)
	}
	productsSKStatusesConn, err := clientManager.GetClientConn(cfg.Cfg.Services.ProductsSKStatuses)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Products SK Statuses service: %w", err)
	}
	statusesConn, err := clientManager.GetClientConn(cfg.Cfg.Services.Statuses)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Statuses service: %w", err)
	}

	return handle.New(ssoConn, appsConn, locationsConn, movementsConn, productionTasksConn, productSKConn, productsSKStatusesConn, statusesConn), nil
}
