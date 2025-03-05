package apps_handle

import (
	sl "backend/pkg/logger"
	app "backend/protos/gen/go/apps"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (s *AppsService) List(c *gin.Context) {
	op := "apps.List"
	sl.Log.Info("Handling request to list apps", slog.String("op", op))

	resp, err := s.client.Apps(context.Background(), &app.GetAppsRequest{})
	if err != nil {
		sl.Log.Error("Error listing apps", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch apps"})
		return
	}

	sl.Log.Info("Successfully fetched apps", slog.Int("count", len(resp.Data)), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
