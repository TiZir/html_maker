package api

import (
	"github.com/TiZir/html_maker/db"
)

type Book interface {
	GetAll() ([]db.BookBody, error)
}

type API struct {
	Book
}

func NewAPI(repos *db.Repository) *API {
	return &API{
		Book: NewBookService(repos.Book),
	}
}
