package universalResponse

import "github.com/gin-gonic/gin"

func SuccessResponse(data interface{}) gin.H {
	return gin.H{
		"success": true,
		"data":    data,
	}
}
