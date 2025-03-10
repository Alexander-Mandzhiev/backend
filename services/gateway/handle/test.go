package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *ServerAPI) healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
