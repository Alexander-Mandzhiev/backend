package location_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/universalResponse"
	"backend/protos/gen/go/locations"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) list(c *gin.Context) {
	op := "locations.List"
	sl.Log.Info("Listing locations", slog.String("op", op))

	resp, err := h.service.List(c.Request.Context(), &locations.ListLocationsRequest{})
	if err != nil {
		sl.Log.Error("List failed", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse("Failed to list locations"))
		return
	}

	sl.Log.Info("Locations listed", slog.Int("count", len(resp.Data)), slog.String("op", op))
	c.JSON(http.StatusOK, universalResponse.SuccessResponse(resp))
}
