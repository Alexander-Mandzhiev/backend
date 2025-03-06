package location_types_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (s *LocationTypesService) Create(c *gin.Context) {
	op := "location_types.Create"
	var req location_types.CreateLocationTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sl.Log.Warn("Invalid request body", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sl.Log.Info("Creating new location type", slog.String("name", req.GetName()), slog.String("description", req.GetDescription()), slog.String("op", op))

	resp, err := s.client.Create(context.Background(), &req)
	if err != nil {
		sl.Log.Error("Error creating location type", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sl.Log.Info("Location type created successfully", slog.Int("id", int(resp.Id)), slog.String("op", op))
	c.JSON(http.StatusCreated, resp)
}
