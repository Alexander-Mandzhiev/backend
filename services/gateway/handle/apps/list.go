package apps_handle

import (
	"backend/pkg/server/universalResponse"
	app "backend/protos/gen/go/apps"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) list(c *gin.Context) {
	op := "apps.List"
	slog.Info("Fetching list of apps", slog.String("op", op))

	resp, err := h.service.List(c.Request.Context(), &app.GetAppsRequest{})
	if err != nil {
		slog.Error("Error listing apps", slog.String("op", op), slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse("Failed to list apps"))
		return
	}

	slog.Info("Apps listed successfully", slog.String("op", op), slog.Int("count", len(resp.Data)))
	c.JSON(http.StatusOK, universalResponse.SuccessResponse(resp))
}
