package repository

import (
	"library/domain/member/model"
	"time"

	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		GetList() ([]model.Member, error)
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

func (r *repository) GetList() ([]model.Member, error) {
	var list []model.Member
	return list, r.DB.Find(&list).Error
}

func (r *repository) Take(selectParams []string, conditions *model.Member) (member model.Member, err error) {
	return member, r.DB.Select(selectParams).Take(&member, conditions).Error
}

func (r *repository) UpdateReturn(code string, date time.Time) error {
	return r.DB.Debug().Table("members").Where("code = ?", code).Update("date_penalty_expired", date).Error
}