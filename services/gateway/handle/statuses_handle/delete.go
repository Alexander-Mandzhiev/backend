package statuses_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *StatusesService) Delete(c *gin.Context) {
	op := "statuses.Delete"
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		sl.Log.Warn("Invalid ID provided", slog.String("id", idStr), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	sl.Log.Info("Deleting status", slog.Int("id", int(id)), slog.String("op", op))

	req := statuses.DeleteStatusRequest{Id: int32(id)}
	resp, err := s.client.Delete(c.Request.Context(), &req)
	if err != nil {
		sl.Log.Error("Error deleting status", sl.Err(err), slog.Int("id", int(id)), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete status"})
		return
	}

	sl.Log.Info("Status deleted successfully", slog.Int("id", int(id)), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
