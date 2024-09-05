package auth

import (
	"app/models/responses"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	response := responses.New(c)
	response.AddClientError("Invalid username or password", 400)
	return response.ClientException()
}
