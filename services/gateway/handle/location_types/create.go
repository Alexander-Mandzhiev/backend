package location_types_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/universalResponse"
	"backend/protos/gen/go/location_types"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	sl.Log.Info("Handling create app request")
	var req location_types.CreateLocationTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sl.Log.Error("Invalid request format", sl.Err(err))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse(err.Error()))
		return
	}

	resp, err := h.service.Create(c, &req)
	if err != nil {
		sl.Log.Error("Failed to create app", sl.Err(err))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse(err.Error()))
		return
	}

	sl.Log.Info("App created successfully", slog.Int("app_id", int(resp.Id)))
	c.JSON(http.StatusOK, universalResponse.SuccessResponse(resp))
}
