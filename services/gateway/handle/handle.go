package handle

import (
	"backend/services/gateway/handle/apps_handle"
	"backend/services/gateway/handle/location_handle"
	"backend/services/gateway/handle/movements_handle"
	"backend/services/gateway/handle/production_task_handle"
	"backend/services/gateway/handle/products_sk_handle"
	"backend/services/gateway/handle/products_sk_statuses_handle"
	"backend/services/gateway/handle/sso_handle"
	"backend/services/gateway/handle/statuses_handle"
	"google.golang.org/grpc"
	"net/http"
)

type Handler struct {
	locationsClient          *location_handle.LocationService
	appsClient               *apps_handle.AppsService
	ssoClient                *sso_handle.SSOService
	movementsClient          *movements_handle.MovementsHandle
	productionTasksClient    *production_task_handle.ProductionTaskHandle
	productSKClient          *products_sk_handle.ProductionSkHandle
	productsSKStatusesClient *products_sk_statuses_handle.ProductsSkStatusesService
	statusesClient           *statuses_handle.StatusesService
}

func New(ssoConn, appsConn, locationsConn, movementsConn, productionTasksConn, productSKConn, productsSKStatusesConn, statusesConn *grpc.ClientConn) *Handler {
	return &Handler{
		ssoClient:                sso_handle.New(ssoConn),
		appsClient:               apps_handle.New(appsConn),
		locationsClient:          location_handle.New(locationsConn),
		movementsClient:          movements_handle.New(movementsConn),
		productionTasksClient:    production_task_handle.New(productionTasksConn),
		productSKClient:          products_sk_handle.New(productSKConn),
		productsSKStatusesClient: products_sk_statuses_handle.New(productsSKStatusesConn),
		statusesClient:           statuses_handle.New(statusesConn),
	}
}

func (h *Handler) InitRouters() http.Handler {
	mux := http.NewServeMux()
	h.initRoutes(mux)
	return mux
}

func (h *Handler) initRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/test", h.test)
	mux.HandleFunc("/api/v1/location/", h.locationsClient.HandleLocation)
}
