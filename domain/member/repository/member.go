package repository

import (
	"library/domain/member/model"
	"time"

	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		GetList() ([]model.List, error)
		Take(selectParams []string, conditions *model.Member) (book model.Member, err error)
		UpdateReturn(code string, date time.Time) error
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

func (r *repository) GetList() ([]model.List, error) {
	var list []model.List

	return list, r.DB.Debug().Model(&model.Member{}).Select("members.name as name, COUNT(books.code) as quantity").Joins("LEFT JOIN books on members.code = books.borrowed_by").Group("members.name").Scan(&list).Error
}

func (r *repository) Take(selectParams []string, conditions *model.Member) (member model.Member, err error) {
	return member, r.DB.Select(selectParams).Take(&member, conditions).Error
}

func (r *repository) UpdateReturn(code string, date time.Time) error {
	return r.DB.Debug().Table("members").Where("code = ?", code).Update("date_penalty_expired", date).Error
}