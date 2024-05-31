package book

import (
	"library/domain/book/model"
	"library/domain/book/repository"
)

type Service struct {
	Repository repository.RepositoryInterface
}

func NewBookService(repository repository.RepositoryInterface) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) List() ([]model.Book, error) {
	data, err := s.Repository.GetList()
	if err != nil {
		return data, err
	}
	return data, nil
}
