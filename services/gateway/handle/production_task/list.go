package production_task_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/universalResponse"
	"backend/protos/gen/go/production_task"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"net/http"
	"time"
)

func (h *Handler) list(c *gin.Context) {
	op := "production_task.List"
	var req struct {
		SklID     int32  `json:"skl_id"`
		DateStart string `json:"date_start" binding:"required"`
		DateEnd   string `json:"date_end" binding:"required"`
		Search    string `json:"search"`
		Page      int32  `json:"page" binding:"min=1"`
		Count     int32  `json:"count" binding:"min=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		sl.Log.Error("Invalid request parameters", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid request parameters"))
		return
	}

	dateStart, err := time.Parse(time.RFC3339, req.DateStart)
	if err != nil {
		sl.Log.Error("Invalid date format", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid date format"))
		return
	}

	dateEnd, err := time.Parse(time.RFC3339, req.DateEnd)
	if err != nil {
		sl.Log.Error("Invalid date format", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid date format"))
		return
	}

	if dateStart.After(dateEnd) {
		sl.Log.Warn("Invalid date range", slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("date_start must be before date_end"))
		return
	}

	reqProto := &production_task.RequestTaskParams{
		SklId:     req.SklID,
		DateStart: timestamppb.New(dateStart),
		DateEnd:   timestamppb.New(dateEnd),
		Search:    req.Search,
		Page:      req.Page,
		Count:     req.Count,
	}

	resp, err := h.service.List(c.Request.Context(), reqProto)
	if err != nil {
		sl.Log.Error("Failed to list tasks", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse("Failed to list tasks"))
		return
	}

	sl.Log.Info("Tasks listed successfully", slog.Int("count", len(resp.GetData())), slog.String("op", op))
	c.JSON(http.StatusOK, universalResponse.SuccessResponse(resp))
}
