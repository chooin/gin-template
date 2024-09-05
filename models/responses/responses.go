package responses

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Responses interface {
	AddClientError(message string, code int)
	HasClientError() bool
	Ok(data any, code ...int) error
	NoContent() error
	ServerException(message string, code ...int) error
	ClientException(code ...int) error
}

type Data = map[string]interface{}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Exception struct {
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
}

type ResponseImpl struct {
	c            echo.Context
	clientErrors []Error
}

func (r *ResponseImpl) Ok(data any, code ...int) error {
	httpStatus := http.StatusOK
	if len(code) > 0 && code[0] >= 200 && code[0] < 300 {
		httpStatus = code[0]
	}
	return r.c.JSON(httpStatus, data)
}

func (r *ResponseImpl) NoContent() error {
	return r.c.NoContent(http.StatusNoContent)
}

func (r *ResponseImpl) ClientException(code ...int) error {
	httpStatus := http.StatusBadRequest
	if len(code) > 0 && code[0] >= 400 && code[0] < 500 {
		httpStatus = code[0]
	}
	return r.c.JSON(httpStatus, &Exception{
		Message: r.clientErrors[0].Message,
		Errors:  r.clientErrors,
	})
}

func (r *ResponseImpl) ServerException(message string, code ...int) error {
	httpStatus := http.StatusInternalServerError
	if len(code) > 0 && code[0] >= 500 && code[0] < 600 {
		httpStatus = code[0]
	}
	return r.c.JSON(httpStatus, Data{
		"message": message,
	})
}

func (r *ResponseImpl) AddClientError(message string, code int) {
	r.clientErrors = append(r.clientErrors, Error{
		Code:    code,
		Message: message,
	})
}

func (r *ResponseImpl) HasClientError() bool {
	return len(r.clientErrors) > 0
}

func New(c echo.Context) Responses {
	return &ResponseImpl{
		c: c,
	}
}
