package statuses_handle

import (
	sl "backend/pkg/logger"
	http_app "backend/pkg/server/universalResponse"
	"backend/protos/gen/go/statuses"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) delete(c *gin.Context) {
	op := "statuses.Delete"
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		sl.Log.Warn("Invalid ID", sl.Err(err), slog.String("id", idStr), slog.String("op", op))
		c.JSON(http.StatusBadRequest, http_app.ErrorResponse("Invalid ID"))
		return
	}

	sl.Log.Info("Deleting status", slog.Int("id", int(id)), slog.String("op", op))

	req := &statuses.DeleteStatusRequest{Id: int32(id)}
	resp, err := h.service.Delete(c.Request.Context(), req)
	if err != nil {
		sl.Log.Error("Delete failed", sl.Err(err), slog.Int("id", int(id)), slog.String("op", op))
		c.JSON(http.StatusNotFound, http_app.ErrorResponse("Status not found"))
		return
	}

	sl.Log.Info("Status deleted", slog.Bool("success", resp.Success), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
