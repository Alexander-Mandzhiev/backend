package location_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *LocationService) Delete(c *gin.Context) {
	op := "location.Delete"
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		sl.Log.Warn("Invalid ID provided", slog.String("id", idStr), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	sl.Log.Info("Deleting location", slog.Int("id", int(id)), slog.String("op", op))

	resp, err := s.client.Delete(context.Background(), &locations.DeleteLocationRequest{Id: int32(id)})
	if err != nil {
		sl.Log.Error("Error deleting location", sl.Err(err), slog.Int("id", int(id)), slog.String("op", op))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	sl.Log.Info("Location deleted successfully", slog.Int("id", int(id)), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
