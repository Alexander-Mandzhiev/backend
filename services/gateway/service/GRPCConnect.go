package service

import (
	"backend/pkg/server/grpc_client"
	cfg "backend/services/gateway/config"

	"fmt"
	"google.golang.org/grpc"
	"log"
)

type GRPCConnections struct {
	SsoConn                *grpc.ClientConn
	AppsConn               *grpc.ClientConn
	LocationsConn          *grpc.ClientConn
	LocationTypesConn      *grpc.ClientConn
	MovementsConn          *grpc.ClientConn
	ProductionTasksConn    *grpc.ClientConn
	ProductSKConn          *grpc.ClientConn
	ProductsSKStatusesConn *grpc.ClientConn
	StatusesConn           *grpc.ClientConn
}

// CreateGRPCClients создает все необходимые gRPC клиенты
func CreateGRPCClients() (*GRPCConnections, error) {
	clientManager := grpc_client.NewGRPCClientManager()
	var conns GRPCConnections
	var err error
	var errs []error

	conns.SsoConn, err = clientManager.GetClientConn(cfg.Cfg.Services.Sso)
	if err != nil {
		errs = append(errs, fmt.Errorf("SSO connection failed: %w", err))
	} else {
		log.Println("Successfully connected to SSO service")
	}

	conns.AppsConn, err = clientManager.GetClientConn(cfg.Cfg.Services.Apps)
	if err != nil {
		errs = append(errs, fmt.Errorf("Apps connection failed: %w", err))
	} else {
		log.Println("Successfully connected to Apps service")
	}

	conns.LocationsConn, err = clientManager.GetClientConn(cfg.Cfg.Services.Locations)
	if err != nil {
		errs = append(errs, fmt.Errorf("Locations connection failed: %w", err))
	} else {
		log.Println("Successfully connected to Locations service")
	}

	conns.LocationTypesConn, err = clientManager.GetClientConn(cfg.Cfg.Services.LocationTypes)
	if err != nil {
		errs = append(errs, fmt.Errorf("Location Types connection failed: %w", err))
	} else {
		log.Println("Successfully connected to Location Types service")
	}

	conns.MovementsConn, err = clientManager.GetClientConn(cfg.Cfg.Services.Movements)
	if err != nil {
		errs = append(errs, fmt.Errorf("Movements connection failed: %w", err))
	} else {
		log.Println("Successfully connected to Movements service")
	}

	conns.ProductionTasksConn, err = clientManager.GetClientConn(cfg.Cfg.Services.ProductionTasks)
	if err != nil {
		errs = append(errs, fmt.Errorf("Production Tasks connection failed: %w", err))
	} else {
		log.Println("Successfully connected to Production Tasks service")
	}

	conns.ProductSKConn, err = clientManager.GetClientConn(cfg.Cfg.Services.ProductSK)
	if err != nil {
		errs = append(errs, fmt.Errorf("Product SK connection failed: %w", err))
	} else {
		log.Println("Successfully connected to Product SK service")
	}

	conns.ProductsSKStatusesConn, err = clientManager.GetClientConn(cfg.Cfg.Services.ProductsSKStatuses)
	if err != nil {
		errs = append(errs, fmt.Errorf("Products SK Statuses connection failed: %w", err))
	} else {
		log.Println("Successfully connected to Products SK Statuses service")
	}

	conns.StatusesConn, err = clientManager.GetClientConn(cfg.Cfg.Services.Statuses)
	if err != nil {
		errs = append(errs, fmt.Errorf("Statuses connection failed: %w", err))
	} else {
		log.Println("Successfully connected to Statuses service")
	}

	if len(errs) > 0 {
		conns.CloseAll()
		return nil, fmt.Errorf("failed to initialize all connections: %v", errs)
	}

	return &conns, nil
}

func (c *GRPCConnections) CloseAll() {
	conns := []*grpc.ClientConn{
		c.SsoConn,
		c.AppsConn,
		c.LocationsConn,
		c.LocationTypesConn,
		c.MovementsConn,
		c.ProductionTasksConn,
		c.ProductSKConn,
		c.ProductsSKStatusesConn,
		c.StatusesConn,
	}

	for _, conn := range conns {
		if conn != nil {
			if err := conn.Close(); err != nil {
				log.Printf("Failed to close gRPC connection: %v", err)
			} else {
				log.Println("gRPC connection closed successfully")
			}
		}
	}
}
