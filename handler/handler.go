package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func HomePage(c echo.Context) error {
	return c.String(http.StatusOK, "HTML MAKER IS RUNNING")
}
