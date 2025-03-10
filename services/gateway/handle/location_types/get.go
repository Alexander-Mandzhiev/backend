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

func (h *Handler) get(c *gin.Context) {
	op := "location_types.Get"
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		sl.Log.Warn("Invalid ID", sl.Err(err), slog.String("id", idStr), slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid ID"))
		return
	}

	req := &location_types.GetLocationTypeRequest{Id: int32(id)}
	resp, err := h.service.Get(c, req)
	if err != nil {
		sl.Log.Error("Failed to get location type", sl.Err(err))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse("Get failed"))
		return
	}

	sl.Log.Info("Location types retrieved", slog.Int("id", int(resp.Id)), slog.String("op", op))
	c.JSON(http.StatusOK, universalResponse.SuccessResponse(resp))
}
