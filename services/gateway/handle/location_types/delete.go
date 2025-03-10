package location_types_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/universalResponse"
	"backend/protos/gen/go/location_types"
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

	req := &location_types.DeleteLocationTypeRequest{Id: int32(id)}
	resp, err := h.service.Delete(c, req)
	if err != nil {
		sl.Log.Error("Failed to delete location type", sl.Err(err))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse("Delete failed"))
		return
	}

	sl.Log.Info("Location types deleted", slog.Bool("success", resp.Success), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
