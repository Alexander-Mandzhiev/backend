package handle

import (
	sl "backend/pkg/logger"
	"backend/services/gateway/handle/apps_handle"
	"backend/services/gateway/handle/location_handle"
	"backend/services/gateway/handle/movements_handle"
	"backend/services/gateway/handle/production_task_handle"
	"backend/services/gateway/handle/products_sk_handle"
	"backend/services/gateway/handle/products_sk_statuses_handle"
	"backend/services/gateway/handle/sso_handle"
	"backend/services/gateway/handle/statuses_handle"
	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: sl.NewLoggerWriter(sl.Log), // Используем нашу обертку
	}))

	api := router.Group("/api/v1")
	router.Use(gin.Recovery())

	h.initLocationRoutes(api)
	h.initAppsRoutes(api)
	h.initStatusesRoutes(api)
	h.initSSORoutes(api)
	return router
}

func (h *Handler) initLocationRoutes(api *gin.RouterGroup) {
	api.GET("/location", h.locationsClient.List)
	api.POST("/location", h.locationsClient.Create)
	api.GET("/location/:id", h.locationsClient.Get)
	api.PUT("/location/:id", h.locationsClient.Update)
	api.DELETE("/location/:id", h.locationsClient.Delete)
}

func (h *Handler) initAppsRoutes(api *gin.RouterGroup) {
	api.POST("/apps", h.appsClient.Create)
	api.GET("/apps", h.appsClient.List)
	api.GET("/apps/:id", h.appsClient.Get)
	api.PUT("/apps/:id", h.appsClient.Update)
	api.DELETE("/apps/:id", h.appsClient.Delete)
}

func (h *Handler) initStatusesRoutes(api *gin.RouterGroup) {
	api.POST("/statuses", h.statusesClient.Create)
	api.GET("/statuses", h.statusesClient.List)
	api.GET("/statuses/:id", h.statusesClient.Get)
	api.PUT("/statuses/:id", h.statusesClient.Update)
	api.DELETE("/statuses/:id", h.statusesClient.Delete)
}

func (h *Handler) initSSORoutes(api *gin.RouterGroup) {
	api.POST("/login", h.ssoClient.SignIn)
}
