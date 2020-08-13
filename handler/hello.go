package handler

import (
	"fmt"
	"hello/server"
	"hello/service"
	"hello/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloHandler struct {
	Service service.Service
	Server  server.Server
}

func (h HelloHandler) InitRoutes() {
	h.Server.Echo.GET("/hello", h.Hello)
}

// Handler
func (h HelloHandler) Hello(c echo.Context) error {
	fmt.Println(h.Server.Echo)
	u := &user.User{
		Name: "Jason",
		ID:   "1234",
	}
	return c.JSON(http.StatusOK, u)
}

func NewHelloHandler(s service.Service, serve server.Server) HelloHandler {
	h := HelloHandler{
		Service: s,
		Server:  serve,
	}
	h.InitRoutes()
	return h
}
