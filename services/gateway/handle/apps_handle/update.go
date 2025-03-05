package apps_handle

import (
	sl "backend/pkg/logger"
	app "backend/protos/gen/go/apps"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (s *AppsService) Update(c *gin.Context) {
	op := "apps.Update"
	var req app.UpdateRequest
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

	sl.Log.Info("Updating app", slog.Int("id", int(req.Id)), slog.String("name", req.GetName()), slog.String("op", op))

	resp, err := s.client.Update(context.Background(), &req)
	if err != nil {
		sl.Log.Error("Error updating app", sl.Err(err), slog.Int("id", int(req.Id)), slog.String("op", "apps.Update"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update app"})
		return
	}

	sl.Log.Info("App updated successfully", slog.Int("id", int(req.Id)), slog.String("op", "apps.Update"))
	c.JSON(http.StatusOK, resp)
}
