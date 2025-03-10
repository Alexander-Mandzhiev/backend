package apps_handle

import (
	"backend/pkg/server/universalResponse"
	app "backend/protos/gen/go/apps"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) get(c *gin.Context) {
	op := "apps.Get"
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		slog.Warn("Invalid ID provided", slog.String("id", idStr), slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid ID"))
		return
	}

	resp, err := h.service.Get(c.Request.Context(), &app.GetAppRequest{Id: int32(id)})
	if err != nil {
		slog.Error("Error fetching app", slog.String("op", op), slog.Any("error", err), slog.Int("id", int(id)))
		c.JSON(http.StatusNotFound, universalResponse.ErrorResponse("App not found"))
		return
	}

	slog.Info("App fetched successfully", slog.Int("id", int(resp.App.Id)), slog.String("op", op))
	c.JSON(http.StatusOK, universalResponse.SuccessResponse(resp))
}
