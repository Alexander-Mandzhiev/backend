package apps_handle

import (
	sl "backend/pkg/logger"
	app "backend/protos/gen/go/apps"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (s *AppsService) Create(c *gin.Context) {
	op := "apps.Create"
	var req app.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sl.Log.Warn("Invalid request body", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	sl.Log.Info("Creating new app", slog.String("name", req.GetName()), slog.String("op", op))

	resp, err := s.client.Create(context.Background(), &req)
	if err != nil {
		sl.Log.Error("Error creating app", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create app"})
		return
	}

	sl.Log.Info("App created successfully", slog.Int("id", int(resp.Id)), slog.String("op", op))
	c.JSON(http.StatusCreated, resp)
}
