package response

import "github.com/gin-gonic/gin"


func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Data: EmptyData,
		Message: message,
	})
}