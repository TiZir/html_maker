package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func GetDB() (*sql.DB, error) {
	return sql.Open("postgres", os.Getenv("PG_URL"))
}

type Book interface {
	GetAllBooks() ([]BookBody, error)
}

type Repository struct {
	Book
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Book: NewBookPostgres(db),
	}
}
