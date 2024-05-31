package member

import (
	"library/domain/member/model"
	"library/domain/member/repository"
)

type Service struct {
	Repository repository.RepositoryInterface
}

func NewMemberService(repository repository.RepositoryInterface) *Service {
	return &Service{
		Repository: repository,
	}
}

func(s *Service) List() ([]model.Member, error) {
	data, err := s.Repository.GetList()
	if err != nil {
		return data, err
	}
	return data, nil
}