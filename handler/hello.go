package handler

import (
	"hello/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler
func Hello(c echo.Context) error {
	u := &user.User{
		Name: "Jason",
		ID:   "1234",
	}
	return c.JSON(http.StatusOK, u)
}
