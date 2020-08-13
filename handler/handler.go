package handler

import (
	"hello/model"
	"hello/server"
	"hello/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service service.Service
	Server  server.Server
}

func (h Handler) InitRoutes() {
	h.Server.Echo.GET("/", h.Hello)
}

func (h Handler) Hello(c echo.Context) error {
	resp, _ := h.Service.Client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetResult(model.Response{}).
		Get("https://run.mocky.io/v3/765cf16a-195c-46cb-a128-2c1789f69037")
	result := resp.Result().(*model.Response)
	return c.JSON(http.StatusOK, result)
}

func NewHandler(s service.Service, serve server.Server) Handler {
	h := Handler{
		Service: s,
		Server:  serve,
	}
	h.InitRoutes()
	return h
}

// var HandlerSet = wire.NewSet(NewHandler, service.NewService)
