package location_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/universalResponse"
	"backend/protos/gen/go/locations"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	op := "locations.Create"
	var req locations.CreateLocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sl.Log.Warn("Invalid request body", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid request body"))
		return
	}

	resp, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		sl.Log.Error("Create failed", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse("Failed to create location"))
		return
	}

	sl.Log.Info("Location created", slog.Int("id", int(resp.Id)), slog.String("op", op))
	c.JSON(http.StatusCreated, universalResponse.SuccessResponse(resp))
}
