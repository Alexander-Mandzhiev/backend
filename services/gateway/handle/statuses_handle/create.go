package statuses_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (s *StatusesService) Create(c *gin.Context) {
	op := "statuses.Create"

	var req statuses.CreateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sl.Log.Warn("Invalid request body", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	sl.Log.Info("Creating new status", slog.String("name", req.GetName()), slog.String("op", op))

	resp, err := s.client.Create(c.Request.Context(), &req)
	if err != nil {
		sl.Log.Error("Error creating status", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create status"})
		return
	}

	sl.Log.Info("Status created successfully", slog.Int("id", int(resp.Id)), slog.String("op", op))
	c.JSON(http.StatusCreated, resp)
}
