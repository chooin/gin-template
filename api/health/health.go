package health

import (
	"app/models/responses"
	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	response := responses.New(c)
	return response.Ok(responses.Data{
		"status": "UP",
	})
}
