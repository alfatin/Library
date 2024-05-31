package book

import (
	"library/domain/book"
	"library/domain/book/model"
	"library/domain/book/repository"
	memberRepo "library/domain/member/repository"
	"library/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type controller struct {
	service     *book.Service
}

func NewController(db *gorm.DB) *controller {
	return &controller{service: book.NewBookService(repository.NewBookRepository(db), memberRepo.NewMemberRepository(db))}
}

// func (c *controller) List(ctx *gin.Context) {
// 	// data, err := c.service.List()
// 	if err != nil {
// 		response.Error(ctx, http.StatusInternalServerError, err.Error())
// 	}
// 	response.Success(ctx, http.StatusOK, data)
// }

func (c *controller) Order(ctx *gin.Context) {
	var req *model.OrderReq
	
	if err := ctx.BindQuery(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	statusCode, err := c.service.Order(req)
	if err != nil {
		response.Error(ctx, statusCode, err.Error())
	}
	response.Success(ctx, statusCode, nil)
}
