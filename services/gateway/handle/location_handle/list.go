package location_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/locations"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (s *LocationService) List(c *gin.Context) {
	op := "location.List"
	sl.Log.Info("Handling request to list locations", slog.String("op", op))

	resp, err := s.client.List(context.Background(), &locations.ListLocationsRequest{})
	if err != nil {
		sl.Log.Error("Error listing locations", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sl.Log.Info("Successfully fetched locations", slog.Int("count", len(resp.Data)), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
