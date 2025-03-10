package location_types_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/universalResponse"
	"backend/protos/gen/go/location_types"
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

	var req location_types.UpdateLocationTypeRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		sl.Log.Error("Invalid request format", sl.Err(err))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid request format"))
		return
	}

	resp, err := h.service.Update(c, &req)
	if err != nil {
		sl.Log.Error("Failed to update location type", sl.Err(err))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse("Update failed"))
		return
	}

	sl.Log.Info("Location updated", slog.Int("id", int(req.Id)), slog.String("op", op))
	c.JSON(http.StatusOK, universalResponse.SuccessResponse(resp))
}
