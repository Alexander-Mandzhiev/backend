package apps_handle

import (
	"backend/pkg/server/universalResponse"
	app "backend/protos/gen/go/apps"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	op := "apps.Create"
	var req app.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Warn("Invalid request body", slog.String("op", op), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid request body"))
		return
	}

	resp, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		slog.Error("Error creating app", slog.String("op", op), slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse("Failed to create app"))
		return
	}

	slog.Info("App created successfully", slog.Int("id", int(resp.Id)), slog.String("op", op))
	c.JSON(http.StatusCreated, universalResponse.SuccessResponse(resp))
}
