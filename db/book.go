package db

import (
	"database/sql"
	"fmt"
)

type BookPostgres struct {
	db *sql.DB
}

func NewBookPostgres(db *sql.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (repo *BookPostgres) GetAllBooks() ([]BookBody, error) {
	var books []BookBody
	query := "SELECT * FROM dev.book"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book BookBody
		err = rows.Scan(&book.ID, &book.Author, &book.Title, &book.Genre, &book.Copies, &book.Rating)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
		fmt.Println("book", book)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
