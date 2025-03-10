package universalResponse

import "github.com/gin-gonic/gin"

func ErrorResponse(errMessage string) gin.H {
	return gin.H{
		"success": false,
		"error":   errMessage,
	}
}
