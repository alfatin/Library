package book

import (
	"errors"
	"library/domain/book/model"
	"library/domain/book/repository"
	memberModel "library/domain/member/model"
	memberRepo "library/domain/member/repository"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	Repository repository.RepositoryInterface
	MemberRepo memberRepo.RepositoryInterface
}

func NewBookService(repository repository.RepositoryInterface, memberRepo memberRepo.RepositoryInterface) *Service {
	return &Service{
		Repository: repository,
		MemberRepo: memberRepo,
	}
}

// func (s *Service) List() ([]model.Book, error) {
// 	data, err := s.Repository.GetList()
// 	if err != nil {
// 		return data, err
// 	}
// 	return data, nil
// }

func (s *Service) Order(req *model.OrderReq) (statusCode int, err error) {
	data, err := s.Repository.Take([]string{"*"}, &model.Book{Code: req.BookCode})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusBadRequest, errors.New("kode buku tidak ditemukan")
		}
		return http.StatusInternalServerError, err
	}

	if data.BorrowedBy != "" && req.Status == "order" {
		return http.StatusBadRequest, errors.New("buku sedang dipinjamkan ke member lain")
	}

	dataMember, err := s.MemberRepo.Take([]string{"*"}, &memberModel.Member{Code: req.MemberCode})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusBadRequest, errors.New("kode member tidak ditemukan")
		}
		return http.StatusInternalServerError, err
	}

	if time.Now().Before(dataMember.DatePenaltyExpired) && req.Status == "order" {
		return http.StatusBadRequest, errors.New("member dalam masa penalti")
	}

	if req.Status == "return" {
		diff := time.Since(data.DateBorrowed)
		days := int(diff.Hours() / 24)

		if days > 7 {
			err = s.MemberRepo.UpdateReturn(req.MemberCode, time.Now().AddDate(0,0,3))
			if err != nil {
				return http.StatusInternalServerError, err
			}
		}
		err = s.Repository.UpdateReturn(req.BookCode)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	if req.Status == "order" {
		dataBook, err := s.Repository.GetList([]string{"code"}, &model.Book{BorrowedBy: req.MemberCode})
		if err != nil {
			return http.StatusInternalServerError, err
		}

		if len(dataBook) > 1 {
			return http.StatusBadRequest, errors.New("member sudah sampai jumlah maksimum peminjaman buku")
		}

		err = s.Repository.UpdateBorrow(req.BookCode, req.MemberCode)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusOK, nil
}
