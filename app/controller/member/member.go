package member

import (
	"library/domain/member"
	"library/domain/member/repository"
	"library/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type controller struct {
	service     *member.Service
}

func NewController(db *gorm.DB) *controller {
	return &controller{service: member.NewMemberService(repository.NewMemberRepository(db))}
}

func (c *controller) List(ctx *gin.Context) {
	data, err := c.service.List()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
	}
	response.Success(ctx, http.StatusOK, data)
}
