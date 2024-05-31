package repository

import (
	"library/domain/book/model"
	"time"

	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		GetList(selectParams []string, conditions *model.Book) ([]model.Book, error)
		GetListExcept() ([]model.Book, error)
		Take(selectParams []string, conditions *model.Book) (book model.Book, err error)
		UpdateReturn(code string) error
		UpdateBorrow(code string, borrowedBy string) error
	}

	repository struct {
		DB *gorm.DB
	}
)

func NewBookRepository(DB *gorm.DB) RepositoryInterface {
	return &repository{
		DB: DB,
	}
}

func (r *repository) GetList(selectParams []string, conditions *model.Book) ([]model.Book, error) {
	var list []model.Book
	return list, r.DB.Debug().Select(selectParams).Find(&list, conditions).Error
}

func (r *repository) GetListExcept() ([]model.Book, error) {
	var list []model.Book
	return list, r.DB.Debug().Select("*").Find(&list, "borrowed_by is null").Error
}

func (r *repository) Take(selectParams []string, conditions *model.Book) (book model.Book, err error) {
	return book, r.DB.Debug().Select(selectParams).Take(&book, conditions).Error
}

func (r *repository) UpdateReturn(code string) error {
	return r.DB.Debug().Table("books").Where("code = ?", code).Updates(map[string]interface{}{"borrowed_by": nil, "date_borrowed": nil}).Error
}

func (r *repository) UpdateBorrow(code string, borrowedBy string) error {
	return r.DB.Debug().Table("books").Where("code = ?", code).Updates(map[string]interface{}{"borrowed_by": borrowedBy, "date_borrowed": time.Now()}).Error
}