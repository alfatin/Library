package root

import (
	"library/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct{}


func NewController() *controller {
	return &controller{}
}

func (c *controller) Index(ctx *gin.Context) {
	response.Success(ctx, http.StatusOK, nil)
}
