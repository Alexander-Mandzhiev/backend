package production_task_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/production_task"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"net/http"
	"time"
)

type ListRequest struct {
	SklID     int       `json:"skl_id,omitempty"`
	DateStart time.Time `json:"date_start" binding:"required"`
	DateEnd   time.Time `json:"date_end" binding:"required"`
	Search    string    `json:"search,omitempty"`
	Page      int       `json:"page" default:"1"`
	Count     int       `json:"count" default:"10"`
}

func (s *ProductionTaskHandle) List(c *gin.Context) {
	op := "production_task.List"
	sl.Log.Info("Handling request to list production tasks", slog.String("op", op))
	var listReq ListRequest
	if err := c.ShouldBindJSON(&listReq); err != nil {
		sl.Log.Error("Error parsing request parameters", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}
	if listReq.DateStart.After(listReq.DateEnd) {
		sl.Log.Warn("Invalid date range", slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "date_start must be before or equal to date_end"})
		return
	}

	reqParams := convertToListRequestParams(&listReq)
	resp, err := s.client.GetTasks(context.Background(), reqParams)
	if err != nil {
		sl.Log.Error("Error fetching production tasks", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch production tasks"})
		return
	}

	sl.Log.Info("Successfully fetched production tasks", slog.Int("count", len(resp.Data)), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}

func convertToListRequestParams(req *ListRequest) *production_task.RequestTaskParams {
	return &production_task.RequestTaskParams{
		SklId:     int32(req.SklID),
		DateStart: timestamppb.New(req.DateStart),
		DateEnd:   timestamppb.New(req.DateEnd),
		Search:    req.Search,
		Page:      int32(req.Page),
		Count:     int32(req.Count),
	}
}
