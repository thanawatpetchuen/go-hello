package route

import (
	"hello/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler
func InitRoutes(e *echo.Echo) {
	e.GET("/", handler.Hello)
	e.POST("/login", handler.Login)

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/home", handler.Home)
}
