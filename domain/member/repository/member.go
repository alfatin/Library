package repository

import (
	"library/domain/member/model"

	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		GetList() ([]model.Member, error)
	}

	repository struct {
		DB *gorm.DB
	}
)

func NewMemberRepository(DB *gorm.DB) RepositoryInterface {
	return &repository{
		DB: DB,
	}
}

func (r *repository) GetList() ([]model.Member, error) {
	var list []model.Member
	return list, r.DB.Find(&list).Error
}