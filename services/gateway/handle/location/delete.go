package location_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/universalResponse"
	"backend/protos/gen/go/locations"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) delete(c *gin.Context) {
	op := "locations.Delete"
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		sl.Log.Warn("Invalid ID", sl.Err(err), slog.String("id", idStr), slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid ID"))
		return
	}

	req := &locations.DeleteLocationRequest{Id: int32(id)}
	resp, err := h.service.Delete(c.Request.Context(), req)
	if err != nil {
		sl.Log.Error("Delete failed", sl.Err(err), slog.Int("id", int(id)), slog.String("op", op))
		c.JSON(http.StatusNotFound, universalResponse.ErrorResponse("Location not found"))
		return
	}

	sl.Log.Info("Location deleted", slog.Bool("success", resp.Success), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
