package location_types_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (s *LocationTypesService) Update(c *gin.Context) {
	op := "location_types.Update"
	var req location_types.UpdateLocationTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sl.Log.Warn("Invalid request body", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Id <= 0 {
		sl.Log.Warn("Invalid ID in request", slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	sl.Log.Info("Updating location type", slog.Int("id", int(req.Id)), slog.String("name", req.GetName()), slog.String("description", req.GetDescription()), slog.String("op", op))

	resp, err := s.client.Update(context.Background(), &req)
	if err != nil {
		sl.Log.Error("Error updating location type", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sl.Log.Info("Location type updated successfully", slog.Int("id", int(resp.Id)), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
