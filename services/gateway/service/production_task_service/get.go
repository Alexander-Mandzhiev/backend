package production_task_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/production_task"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (s *ProductionTaskHandle) Get(c *gin.Context) {
	op := "production_task.Get"
	sl.Log.Info("Handling request to get tasks by part name", slog.String("op", op))
	var reqParams production_task.RequestTaskParams
	if err := c.ShouldBindJSON(&reqParams); err != nil {
		sl.Log.Error("Error parsing request parameters", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}

	if reqParams.PartName == "" || reqParams.Page <= 0 || reqParams.Count <= 0 {
		sl.Log.Warn("Missing or invalid required parameters", slog.String("op", op))
		c.JSON(http.StatusBadRequest, gin.H{"error": "part_name, page, and count are required"})
		return
	}

	resp, err := s.client.GetTasksInPartName(context.Background(), &reqParams)
	if err != nil {
		sl.Log.Error("Error fetching tasks by part name", sl.Err(err), slog.String("op", op))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	sl.Log.Info("Successfully fetched tasks by part name", slog.Int("count", len(resp.Data)), slog.String("op", op))
	c.JSON(http.StatusOK, resp)
}
