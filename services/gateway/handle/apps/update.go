package apps_handle

import (
	"backend/pkg/server/universalResponse"
	app "backend/protos/gen/go/apps"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) update(c *gin.Context) {
	op := "apps.Update"
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		slog.Warn("Invalid ID provided", slog.String("id", idStr), slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid ID"))
		return
	}

	var req app.UpdateRequest
	req.Id = int32(id)

	if err = c.ShouldBindJSON(&req); err != nil {
		slog.Warn("Invalid request body", slog.String("op", op), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid request body"))
		return
	}

	resp, err := h.service.Update(c.Request.Context(), &req)
	if err != nil {
		slog.Error("Error updating app", slog.String("op", op), slog.Any("error", err), slog.Int("id", int(req.Id)))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse("Failed to update app"))
		return
	}

	slog.Info("App updated successfully", slog.Int("id", int(req.Id)), slog.String("op", op))
	c.JSON(http.StatusOK, universalResponse.SuccessResponse(resp))
}
