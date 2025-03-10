package location_types_handle

import (
	"backend/protos/gen/go/location_types"
	"context"
	"github.com/gin-gonic/gin"
)

type LocationTypesService interface {
	Create(ctx context.Context, req *location_types.CreateLocationTypeRequest) (*location_types.LocationTypeResponse, error)
	Get(ctx context.Context, req *location_types.GetLocationTypeRequest) (*location_types.LocationTypeResponse, error)
	Update(ctx context.Context, req *location_types.UpdateLocationTypeRequest) (*location_types.LocationTypeResponse, error)
	Delete(ctx context.Context, req *location_types.DeleteLocationTypeRequest) (*location_types.DeleteLocationTypeResponse, error)
	List(ctx context.Context, req *location_types.ListLocationTypesRequest) (*location_types.LocationTypeListResponse, error)
}

type Handler struct {
	service LocationTypesService
}

func New(service LocationTypesService) *Handler {
	return &Handler{service: service}
}
func (h *Handler) InitLocationTypesRoutes(api *gin.RouterGroup) {
	locationTypes := api.Group("/location_types")
	{
		locationTypes.POST("", h.create)
		locationTypes.GET("/:id", h.get)
		locationTypes.PUT("/:id", h.update)
		locationTypes.DELETE("/:id", h.delete)
		locationTypes.GET("", h.list)
	}
}
