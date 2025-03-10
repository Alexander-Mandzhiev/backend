package location_types_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/universalResponse"
	"backend/protos/gen/go/location_types"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) list(c *gin.Context) {
	op := "locations.List"
	sl.Log.Info("Listing location types", slog.String("op", op))

	resp, err := h.service.List(c, &location_types.ListLocationTypesRequest{})
	if err != nil {
		sl.Log.Error("Failed to list location types", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse("List location failed"))
		return
	}

	sl.Log.Info("Location types listed", slog.Int("count", len(resp.Data)), slog.String("op", op))
	c.JSON(http.StatusOK, universalResponse.SuccessResponse(resp))
}
