package statuses_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *StatusesService) Get(c *gin.Context) {
	op := "statuses.Get"
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		sl.Log.Warn("Invalid ID provided", slog.String("id", idStr), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	sl.Log.Info("Fetching status", slog.Int("id", int(id)), slog.String("op", op))

	req := statuses.GetStatusRequest{Id: int32(id)}
	resp, err := s.client.Status(c.Request.Context(), &req)
	if err != nil {
		sl.Log.Error("Error fetching status", sl.Err(err), slog.Int("id", int(id)), slog.String("op", op))
		c.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
		return
	}

	sl.Log.Info("Status fetched successfully", slog.Int("id", int(resp.Id)), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
