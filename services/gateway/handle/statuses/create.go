package statuses_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/universalResponse"
	"backend/protos/gen/go/statuses"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	op := "statuses.Create"
	var req statuses.CreateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sl.Log.Warn("Invalid request body", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid request body"))
		return
	}

	resp, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		sl.Log.Error("Create failed", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse("Failed to create status"))
		return
	}

	sl.Log.Info("Status created", slog.Int("id", int(resp.Id)), slog.String("op", op))
	c.JSON(http.StatusCreated, universalResponse.SuccessResponse(resp))
}
