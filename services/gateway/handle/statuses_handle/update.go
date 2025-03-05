package statuses_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (s *StatusesService) Update(c *gin.Context) {
	op := "statuses.Update"
	var req statuses.UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sl.Log.Warn("Invalid request body", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if req.Id <= 0 {
		sl.Log.Warn("ID must be provided in the request body", slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be provided in the request body"})
		return
	}

	sl.Log.Info("Updating status", slog.Int("id", int(req.Id)), slog.String("name", req.GetName()), slog.String("op", op))

	resp, err := s.client.Update(c.Request.Context(), &req)
	if err != nil {
		sl.Log.Error("Error updating status", sl.Err(err), slog.Int("id", int(req.Id)), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
		return
	}

	sl.Log.Info("Status updated successfully", slog.Int("id", int(req.Id)), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
