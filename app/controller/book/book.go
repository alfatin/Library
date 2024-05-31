package book

import (
	"library/domain/book"
	"library/domain/book/repository"
	"library/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type controller struct {
	service     *book.Service
}

func NewController(db *gorm.DB) *controller {
	return &controller{service: book.NewBookService(repository.NewBookRepository(db))}
}

func (c *controller) List(ctx *gin.Context) {
	data, err := c.service.List()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
	}
	response.Success(ctx, http.StatusOK, data)
}
