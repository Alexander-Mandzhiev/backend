package location_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (s *LocationService) Update(c *gin.Context) {
	op := "location.Update"
	var req locations.UpdateLocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sl.Log.Warn("Invalid request body", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Id <= 0 {
		sl.Log.Warn("ID must be provided in the request body", slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be provided in the request body"})
		return
	}

	sl.Log.Info("Updating location", slog.String("name", req.GetName()), slog.Int("type_id", int(req.GetTypeId())), slog.Int("capacity", int(req.GetCapacity())), slog.String("op", op))

	resp, err := s.client.Update(context.Background(), &req)
	if err != nil {
		sl.Log.Error("Error updating location", sl.Err(err), slog.Int("id", int(req.Id)), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
