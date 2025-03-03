package handle

import (
	app "backend/protos/gen/go/apps"
	"backend/protos/gen/go/locations"
	"backend/protos/gen/go/movements"
	"backend/protos/gen/go/production_task"
	"backend/protos/gen/go/products_sk"
	"backend/protos/gen/go/products_sk_statuses"
	"backend/protos/gen/go/sso"
	"backend/protos/gen/go/statuses"
	"google.golang.org/grpc"
	"net/http"
)

type Handler struct {
	ssoClient                sso.SSOServiceClient
	appsClient               app.AppProviderServiceClient
	locationsClient          locations.LocationServiceClient
	movementsClient          movements.MovementServiceClient
	productionTasksClient    production_task.ProductionTaskServiceClient
	productSKClient          products_sk.ProductServiceClient
	productsSKStatusesClient products_sk_statuses.ProductStatusServiceClient
	statusesClient           statuses.StatusServiceClient
}

func New(ssoConn, appsConn, locationsConn, movementsConn, productionTasksConn, productSKConn, productsSKStatusesConn, statusesConn *grpc.ClientConn) *Handler {
	return &Handler{
		ssoClient:                sso.NewSSOServiceClient(ssoConn),
		appsClient:               app.NewAppProviderServiceClient(appsConn),
		locationsClient:          locations.NewLocationServiceClient(locationsConn),
		movementsClient:          movements.NewMovementServiceClient(movementsConn),
		productionTasksClient:    production_task.NewProductionTaskServiceClient(productionTasksConn),
		productSKClient:          products_sk.NewProductServiceClient(productSKConn),
		productsSKStatusesClient: products_sk_statuses.NewProductStatusServiceClient(productsSKStatusesConn),
		statusesClient:           statuses.NewStatusServiceClient(statusesConn),
	}
}
func (h *Handler) InitRouters() http.Handler {
	mux := http.NewServeMux()
	h.initRoutes(mux)
	return mux
}

func (h *Handler) initRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/test", h.test)
	/*mux.HandleFunc("/auth/login", h)
	mux.HandleFunc("/auth/signup", h.signUp)
	mux.HandleFunc("/auth/logout", h.logout)
	mux.HandleFunc("/auth/refresh_token", h.refreshToken)

	// Ручки для Goods (защищенные токеном)
	mux.HandleFunc("/goods", h.requireAuth(h.goods))

	// Ручки для Contractor (защищенные токеном)
	mux.HandleFunc("/contractor", h.requireAuth(h.contractor))
	mux.HandleFunc("/statement_contractor", h.requireAuth(h.statementContractor))*/
}
