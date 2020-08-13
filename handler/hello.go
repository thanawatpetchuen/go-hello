package handler

import (
	"hello/service"
	"hello/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloHandler struct {
	Service service.Service
}

// Handler
func (h HelloHandler) Hello(c echo.Context) error {
	u := &user.User{
		Name: "Jason",
		ID:   "1234",
	}
	return c.JSON(http.StatusOK, u)
}

func NewHelloHandler(s service.Service) HelloHandler {
	return HelloHandler{
		Service: s,
	}
}
