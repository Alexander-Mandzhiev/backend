package handle

import (
	sl "backend/pkg/logger"
	"backend/services/gateway/handle/apps"
	"backend/services/gateway/handle/location"
	"backend/services/gateway/handle/location_types"
	"backend/services/gateway/handle/sso"
	"backend/services/gateway/handle/statuses"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServerAPI struct {
	appsClient          *apps_handle.Handler
	ssoClient           *sso_handle.Handler
	statusesClient      *statuses_handle.Handler
	locationsClient     *location_handle.Handler
	locationTypesClient *location_types_handle.Handler
}

func New(appsClient *apps_handle.Handler, ssoClient *sso_handle.Handler, statusesClient *statuses_handle.Handler, locationsClient *location_handle.Handler,
	locationTypesClient *location_types_handle.Handler) *ServerAPI {
	return &ServerAPI{appsClient: appsClient, ssoClient: ssoClient, statusesClient: statusesClient, locationsClient: locationsClient,
		locationTypesClient: locationTypesClient}
}

func (h *ServerAPI) InitRouters() http.Handler {
	router := gin.Default()
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: sl.NewLoggerWriter(sl.Log),
	}))

	router.Use(gin.Recovery())

	api := router.Group("/api/v1") // Группа API v1
	{
		h.appsClient.InitAppsRoutes(api)                   // Инициализация роутов для приложений
		h.ssoClient.InitSSORoutes(api)                     // Роуты аутентификации
		h.statusesClient.InitStatusesRoutes(api)           // Роуты статусов
		h.locationsClient.InitLocationRoutes(api)          // Роуты локаций
		h.locationTypesClient.InitLocationTypesRoutes(api) // Роуты типов локаций
	}

	router.GET("/healthcheck", h.healthcheck)

	return router
}
