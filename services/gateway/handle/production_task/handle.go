package production_task_handle

import (
	"backend/protos/gen/go/production_task"
	"context"
	"github.com/gin-gonic/gin"
)

type ProductionTaskService interface {
	GetTasksByPartName(ctx context.Context, req *production_task.RequestTaskParams) (*production_task.ProductsResponse, error)
	List(ctx context.Context, req *production_task.RequestTaskParams) (*production_task.ProductsResponse, error)
}

type Handler struct {
	service ProductionTaskService
}

func New(service ProductionTaskService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitProductionTaskRoutes(api *gin.RouterGroup) {
	taskRoutes := api.Group("/production_task")
	{
		taskRoutes.POST("/list", h.list)
		taskRoutes.POST("/get", h.getByPartName)
	}
}
