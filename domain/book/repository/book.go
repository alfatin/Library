package repository

import (
	"library/domain/book/model"

	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		GetList() ([]model.Book, error)
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

func (r *repository) GetList() ([]model.Book, error) {
	var list []model.Book
	return list, r.DB.Find(&list).Error
}