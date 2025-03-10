package apps_handle

import (
	app "backend/protos/gen/go/apps"
	"context"
	"github.com/gin-gonic/gin"
)

type AppsService interface {
	Create(ctx context.Context, req *app.CreateRequest) (*app.CreateResponse, error)
	Update(ctx context.Context, req *app.UpdateRequest) (*app.UpdateResponse, error)
	Get(ctx context.Context, req *app.GetAppRequest) (*app.GetAppResponse, error)
	List(ctx context.Context, req *app.GetAppsRequest) (*app.GetAppsResponse, error)
	Delete(ctx context.Context, req *app.DeleteRequest) (*app.DeleteResponse, error)
}

type Handler struct {
	service AppsService
}

func New(service AppsService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitAppsRoutes(api *gin.RouterGroup) {
	appRoutes := api.Group("/apps")

	appRoutes.POST("", h.create)       // Создание приложения
	appRoutes.PUT("/:id", h.update)    // Обновление приложения
	appRoutes.GET("", h.list)          // Получение списка приложений
	appRoutes.GET("/:id", h.get)       // Получение конкретного приложения
	appRoutes.DELETE("/:id", h.delete) // Удаление приложения
}
