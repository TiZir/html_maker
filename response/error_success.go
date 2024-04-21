package response

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type ResponseError struct {
	StatusCode int    `json:"code"`
	Message    string `json:"error"`
}

func (e *ResponseError) Error() string {
	return e.Message
}

func InternalServerErrorResponse(c echo.Context, err error) error {
	rError := ResponseError{StatusCode: http.StatusInternalServerError, Message: "Внутренняя ошибка сервера"}
	return ErrorResponse(c, err, rError)
}

func NotFoundResponse(c echo.Context, err error) error {
	rError := ResponseError{StatusCode: http.StatusNotFound, Message: "Баннер не найден"}
	return ErrorResponse(c, err, rError)
}

func ForbiddenResponse(c echo.Context, err error) error {
	rError := ResponseError{StatusCode: http.StatusForbidden, Message: "Пользователь не имеет доступа"}
	return ErrorResponse(c, err, rError)
}

func UnauthorizedResponse(c echo.Context, err error) error {
	rError := ResponseError{StatusCode: http.StatusUnauthorized, Message: "Пользователь не авторизован"}
	return ErrorResponse(c, err, rError)
}

func BadRequestResponse(c echo.Context, err error) error {
	rError := ResponseError{StatusCode: http.StatusBadRequest, Message: "Некорректные данные"}
	return ErrorResponse(c, err, rError)
}

func ErrorResponse(c echo.Context, err error, rError ResponseError) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(rError.StatusCode)
	errorMsg := err.Error()
	if rError.Message != "" {
		if err.Error() == "" {
			errorMsg = rError.Message
		} else {
			errorMsg = fmt.Sprintf("%s: %s", rError.Message, err.Error())
		}
	}
	return c.JSON(rError.StatusCode, map[string]string{"error": errorMsg})
}

func HandlerError(c echo.Context, err error, statusCode int) error {
	switch statusCode {
	case http.StatusInternalServerError:
		return InternalServerErrorResponse(c, err)
	case http.StatusNotFound:
		return NotFoundResponse(c, err)
	case http.StatusForbidden:
		return ForbiddenResponse(c, err)
	case http.StatusUnauthorized:
		return UnauthorizedResponse(c, err)
	case http.StatusBadRequest:
		return BadRequestResponse(c, err)
	default:
		rError := ResponseError{StatusCode: statusCode, Message: ""}
		return ErrorResponse(c, err, rError)
	}
}
