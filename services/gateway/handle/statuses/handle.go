package statuses_handle

import (
	"backend/protos/gen/go/statuses"
	"context"
	"github.com/gin-gonic/gin"
)

type StatusesService interface {
	Create(ctx context.Context, req *statuses.CreateStatusRequest) (*statuses.StatusResponse, error)
	Update(ctx context.Context, req *statuses.UpdateStatusRequest) (*statuses.StatusResponse, error)
	Get(ctx context.Context, req *statuses.GetStatusRequest) (*statuses.StatusResponse, error)
	List(ctx context.Context, req *statuses.ListStatusesRequest) (*statuses.StatusListResponse, error)
	Delete(ctx context.Context, req *statuses.DeleteStatusRequest) (*statuses.DeleteStatusResponse, error)
}

type Handler struct {
	service StatusesService
}

func New(service StatusesService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitStatusesRoutes(api *gin.RouterGroup) {
	statusRoutes := api.Group("/statuses")

	statusRoutes.POST("", h.create)
	statusRoutes.GET("", h.list)
	statusRoutes.GET("/:id", h.get)
	statusRoutes.PUT("/:id", h.update)
	statusRoutes.DELETE("/:id", h.delete)
}
