package users

import (
	"app/models"
	"app/models/responses"
	"app/pkg/mysql"
	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	response := responses.New(c)
	userId := c.Param("id")
	var user models.User
	mysql.Client.Model(&models.User{}).Where("id = ?", userId).First(&user)
	return response.Ok(user)
}
