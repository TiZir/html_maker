package handler

import (
	"fmt"
	"net/http"

	"github.com/TiZir/html_maker/api"
	"github.com/TiZir/html_maker/response"
	"github.com/labstack/echo"
)

func GetBooks(c echo.Context, service *api.API) error {
	books, err := service.Book.GetAll()
	if err != nil {
		return response.HandlerError(c, err, http.StatusInternalServerError)
	}
	html := "<html><head><title>Books</title></head><body><h1>Books</h1><table><tr><th>ID</th><th>Author</th><th>Title</th><th>Genre</th><th>Copies</th><th>Rating</th></tr>"
	for _, book := range books {
		html += fmt.Sprintf("<tr><td>%d</td><td>%s</td><td>%s</td><td>%s</td><td>%d</td><td>%.1f</td></tr>", book.ID, book.Author, book.Title, book.Genre, book.Copies, book.Rating)
	}
	html += "</table></body></html>"
	return c.HTMLBlob(http.StatusOK, []byte(html))
}
