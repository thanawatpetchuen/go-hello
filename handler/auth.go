package handler

import (
	"hello/auth"
	"hello/model"
	"hello/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	u := user.User{
		Name:     username,
		Password: password,
	}

	if u.Name != "admin" && u.Password != "1234" {
		return echo.ErrUnauthorized
	}

	t, _ := auth.CreateToken(u)

	res := model.LoginResponse{
		Token: t,
	}

	return c.JSON(http.StatusOK, res)
}
