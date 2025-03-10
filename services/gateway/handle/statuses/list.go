package statuses_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/universalResponse"
	"backend/protos/gen/go/statuses"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) list(c *gin.Context) {
	op := "statuses.List"
	sl.Log.Info("Listing statuses", slog.String("op", op))

	resp, err := h.service.List(c.Request.Context(), &statuses.ListStatusesRequest{})
	if err != nil {
		sl.Log.Error("List failed", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse("Failed to list statuses"))
		return
	}

	sl.Log.Info("Statuses listed", slog.Int("count", len(resp.Data)), slog.String("op", op))
	c.JSON(http.StatusOK, universalResponse.SuccessResponse(resp))
}
