package http_app

import (
	"backend/pkg/server/grpc_client"
	cfg "backend/services/gateway/config"
	"log"

	"backend/services/gateway/handle"
	"fmt"
)

func initializeHandler() (*handle.Handler, error) {
	clientManager := grpc_client.NewGRPCClientManager()

	var errs []error

	ssoConn, err := clientManager.GetClientConn(cfg.Cfg.Services.Sso)
	if err != nil {
		log.Printf("Failed to connect to SSO service: %v", err)
		errs = append(errs, fmt.Errorf("SSO connection failed: %w", err))
	}

	appsConn, err := clientManager.GetClientConn(cfg.Cfg.Services.Apps)
	if err != nil {
		log.Printf("Failed to connect to Apps service: %v", err)
		errs = append(errs, fmt.Errorf("apps connection failed: %w", err))
	}

	locationsConn, err := clientManager.GetClientConn(cfg.Cfg.Services.Locations)
	if err != nil {
		log.Printf("Failed to connect to Locations service: %v", err)
		errs = append(errs, fmt.Errorf("locations connection failed: %w", err))
	}
	locationTypesConn, err := clientManager.GetClientConn(cfg.Cfg.Services.LocationTypes)
	if err != nil {
		log.Printf("Failed to connect to Locations service: %v", err)
		errs = append(errs, fmt.Errorf("locations connection failed: %w", err))
	}

	movementsConn, err := clientManager.GetClientConn(cfg.Cfg.Services.Movements)
	if err != nil {
		log.Printf("Failed to connect to Movements service: %v", err)
		errs = append(errs, fmt.Errorf("movements connection failed: %w", err))
	}
	productionTasksConn, err := clientManager.GetClientConn(cfg.Cfg.Services.ProductionTasks)
	if err != nil {
		log.Printf("Failed to connect to Production service: %v", err)
		errs = append(errs, fmt.Errorf("production connection failed: %w", err))
	}
	productSKConn, err := clientManager.GetClientConn(cfg.Cfg.Services.ProductSK)
	if err != nil {
		log.Printf("Failed to connect to Product service: %v", err)
		errs = append(errs, fmt.Errorf("product connection failed: %w", err))
	}
	productsSKStatusesConn, err := clientManager.GetClientConn(cfg.Cfg.Services.ProductsSKStatuses)
	if err != nil {
		log.Printf("Failed to connect to Products sk service: %v", err)
		errs = append(errs, fmt.Errorf("products sk connection failed: %w", err))
	}
	statusesConn, err := clientManager.GetClientConn(cfg.Cfg.Services.Statuses)
	if err != nil {
		log.Printf("Failed to connect to Statuses service: %v", err)
		errs = append(errs, fmt.Errorf("statuses connection failed: %w", err))
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("failed to initialize all connections: %v", errs)
	}

	return handle.New(ssoConn, appsConn, locationsConn, locationTypesConn, movementsConn, productionTasksConn, productSKConn, productsSKStatusesConn, statusesConn), nil
}
