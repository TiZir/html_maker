package api

import (
	"github.com/TiZir/html_maker/db"
)

type BookService struct {
	repo db.Book
}

func NewBookService(repo db.Book) *BookService {
	return &BookService{repo: repo}
}
func (s *BookService) GetAll() ([]db.BookBody, error) {
	return s.repo.GetAllBooks()
}
