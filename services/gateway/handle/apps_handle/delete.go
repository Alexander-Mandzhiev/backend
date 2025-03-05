package apps_handle

import (
	sl "backend/pkg/logger"
	app "backend/protos/gen/go/apps"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *AppsService) Delete(c *gin.Context) {
	op := "apps.Delete"
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		sl.Log.Warn("Invalid ID provided", slog.String("id", idStr), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	sl.Log.Info("Deleting app", slog.Int("id", int(id)), slog.String("op", op))

	resp, err := s.client.Delete(context.Background(), &app.DeleteRequest{Id: int32(id)})
	if err != nil {
		sl.Log.Error("Error deleting app", sl.Err(err), slog.Int("id", int(id)), slog.String("op", op))
		c.JSON(http.StatusNotFound, gin.H{"error": "App not found"})
		return
	}

	sl.Log.Info("App deleted successfully", slog.Int("id", int(id)), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
