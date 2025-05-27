package production_task_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/server/universalResponse"
	"backend/protos/gen/go/production_task"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) getByPartName(c *gin.Context) {
	op := "ProductionTask.GetByPartName"
	var reqData struct {
		PartName string `json:"part_name"`
		SklID    int32  `json:"skl_id"`
		Page     int32  `json:"page"`
		Count    int32  `json:"count"`
	}

	if err := c.ShouldBindJSON(&reqData); err != nil {
		sl.Log.Warn("Invalid request body", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Invalid request parameters"))
		return
	}

	if reqData.PartName == "" {
		sl.Log.Warn("Missing part_name parameter", slog.String("op", op))
		c.JSON(http.StatusBadRequest, universalResponse.ErrorResponse("Part name is required"))
		return
	}

	grpcReq := &production_task.RequestTaskParams{
		PartName: reqData.PartName,
		Page:     reqData.Page,
		Count:    reqData.Count,
		SklId:    reqData.SklID,
	}

	resp, err := h.service.GetTasksByPartName(c.Request.Context(), grpcReq)
	if err != nil {
		sl.Log.Error("Failed to get tasks", sl.Err(err), slog.String("part_name", reqData.PartName), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, universalResponse.ErrorResponse("Failed to retrieve tasks"))
		return
	}

	sl.Log.Info("Tasks retrieved successfully", slog.Int("count", len(resp.GetData())), slog.String("op", op))
	c.JSON(http.StatusOK, universalResponse.SuccessResponse(resp))
}
