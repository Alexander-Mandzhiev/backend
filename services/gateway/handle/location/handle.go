package location_handle

import (
	"backend/protos/gen/go/locations"
	"context"
	"github.com/gin-gonic/gin"
)

type LocationsService interface {
	Create(ctx context.Context, req *locations.CreateLocationRequest) (*locations.LocationResponse, error)
	Update(ctx context.Context, req *locations.UpdateLocationRequest) (*locations.LocationResponse, error)
	Get(ctx context.Context, req *locations.GetLocationRequest) (*locations.LocationResponse, error)
	List(ctx context.Context, req *locations.ListLocationsRequest) (*locations.LocationListResponse, error)
	Delete(ctx context.Context, req *locations.DeleteLocationRequest) (*locations.DeleteLocationResponse, error)
}

type Handler struct {
	service LocationsService
}

func New(service LocationsService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitLocationRoutes(api *gin.RouterGroup) {
	locationRoutes := api.Group("/locations")

	locationRoutes.POST("", h.create)
	locationRoutes.PUT("/:id", h.update)
	locationRoutes.GET("", h.list)
	locationRoutes.GET("/:id", h.get)
	locationRoutes.DELETE("/:id", h.delete)
}
