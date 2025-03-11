package service

import (
	"backend/services/gateway/service/apps_service"
	"backend/services/gateway/service/location_service"
	"backend/services/gateway/service/location_types_service"
	"backend/services/gateway/service/movements_service"
	"backend/services/gateway/service/production_task_service"
	"backend/services/gateway/service/products_sk_service"
	"backend/services/gateway/service/products_sk_statuses_service"
	"backend/services/gateway/service/sso_service"
	"backend/services/gateway/service/statuses_service"
	"google.golang.org/grpc"
)

type Service struct {
	AppsClient *apps_service.Service
	SSOClient  *sso_service.Service

	LocationsClient     *location_service.Service
	LocationTypesClient *location_types_service.Service
	StatusesClient      *statuses_service.Service

	ProductSKClient          *products_sk_service.Service
	MovementsClient          *movements_service.Service
	ProductionTasksClient    *production_task_service.Service
	ProductsSKStatusesClient *products_sk_statuses_service.Service
}

func New(ssoConn, appsConn, locationsConn, locationTypesConn, movementsConn, productionTasksConn, productSKConn,
	productsSKStatusesConn, statusesConn *grpc.ClientConn) *Service {
	return &Service{
		SSOClient:                sso_service.New(ssoConn),
		AppsClient:               apps_service.New(appsConn),
		LocationsClient:          location_service.New(locationsConn),
		LocationTypesClient:      location_types_service.New(locationTypesConn),
		MovementsClient:          movements_service.New(movementsConn),
		ProductionTasksClient:    production_task_service.New(productionTasksConn),
		ProductSKClient:          products_sk_service.New(productSKConn),
		ProductsSKStatusesClient: products_sk_statuses_service.New(productsSKStatusesConn),
		StatusesClient:           statuses_service.New(statusesConn),
	}
}
