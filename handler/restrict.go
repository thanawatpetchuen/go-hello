package handler

import (
	"hello/server"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RestrictHandler struct {
	Server server.Server
}

func (res RestrictHandler) InitRoutes() {
	r := res.Server.Echo.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/home", res.Home)
}

func (res RestrictHandler) Home(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.JSON(http.StatusOK, name)
}

func NewRestrictHandler(s server.Server) RestrictHandler {
	nR := RestrictHandler{s}
	nR.InitRoutes()
	return nR
}
