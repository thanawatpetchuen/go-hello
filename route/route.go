package route

import (
	"hello/handler"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler
func InitRoutes(e *echo.Echo, h handler.Handler) {
	e.GET("/", h.Hello)
	e.POST("/login", handler.Login)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/home", handler.Home)
}
