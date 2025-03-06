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

func (s *LocationTypesService) Delete(c *gin.Context) {
	op := "location_types.Delete"
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		sl.Log.Warn("Invalid ID parameter", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	sl.Log.Info("Deleting location type", slog.Int("id", id), slog.String("op", op))

	req := location_types.DeleteLocationTypeRequest{Id: int32(id)}
	resp, err := s.client.Delete(context.Background(), &req)
	if err != nil {
		sl.Log.Error("Error deleting location type", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sl.Log.Info("Location type deleted successfully", slog.Int("id", id), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
