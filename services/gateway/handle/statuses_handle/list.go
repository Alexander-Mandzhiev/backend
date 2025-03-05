package statuses_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (s *StatusesService) List(c *gin.Context) {
	op := "statuses.List"

	sl.Log.Info("Fetching statuses", slog.String("op", op))

	req := statuses.ListStatusesRequest{}
	resp, err := s.client.List(c.Request.Context(), &req)
	if err != nil {
		sl.Log.Error("Error listing statuses", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch statuses"})
		return
	}

	sl.Log.Info("Statuses fetched successfully", slog.Int("count", len(resp.Data)), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
