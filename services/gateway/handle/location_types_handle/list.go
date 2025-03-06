package location_types_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (s *LocationTypesService) List(c *gin.Context) {
	op := "location_types.List"

	sl.Log.Info("Fetching all location types", slog.String("op", op))

	req := location_types.ListLocationTypesRequest{}
	resp, err := s.client.List(context.Background(), &req)
	if err != nil {
		sl.Log.Error("Error fetching location types", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sl.Log.Info("Location types fetched successfully", slog.Int("count", len(resp.GetData())), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
