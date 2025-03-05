package location_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (s *LocationService) Create(c *gin.Context) {
	op := "location.Create"
	var req locations.CreateLocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sl.Log.Warn("Invalid request body", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sl.Log.Info("Creating new location", slog.String("name", req.GetName()), slog.String("type", req.GetType()), slog.Int("capacity", int(req.GetCapacity())), slog.String("op", op))

	resp, err := s.client.Create(context.Background(), &req)
	if err != nil {
		sl.Log.Error("Error creating location", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sl.Log.Info("Location created successfully", slog.Int("id", int(resp.Id)), slog.String("op", op))
	c.JSON(http.StatusCreated, resp)
}
