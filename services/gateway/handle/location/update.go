package location_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/universalResponse"
	"backend/protos/gen/go/locations"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) update(c *gin.Context) {
	op := "locations.Update"
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		sl.Log.Warn("Invalid ID", sl.Err(err), slog.String("id", idStr), slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid ID"))
		return
	}

	var req locations.UpdateLocationRequest
	req.Id = int32(id)
	if err = c.ShouldBindJSON(&req); err != nil {
		sl.Log.Warn("Invalid request body", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid request body"))
		return
	}

	resp, err := h.service.Update(c.Request.Context(), &req)
	if err != nil {
		sl.Log.Error("Update failed", sl.Err(err), slog.Int("id", int(req.Id)), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse("Failed to update location"))
		return
	}

	sl.Log.Info("Location updated", slog.Int("id", int(req.Id)), slog.String("op", op))
	c.JSON(http.StatusOK, universalResponse.SuccessResponse(resp))
}
