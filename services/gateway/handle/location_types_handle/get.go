package location_types_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/location_types"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *LocationTypesService) Get(c *gin.Context) {
	op := "location_types.Get"
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		sl.Log.Warn("Invalid ID parameter", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	sl.Log.Info("Fetching location type by ID", slog.Int("id", id), slog.String("op", op))

	req := location_types.GetLocationTypeRequest{Id: int32(id)}
	resp, err := s.client.Get(context.Background(), &req)
	if err != nil {
		sl.Log.Error("Error fetching location type", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sl.Log.Info("Location type fetched successfully", slog.Int("id", int(resp.Id)), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
