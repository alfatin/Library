package response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Data: data,
		Message: "OK",
	})
}
