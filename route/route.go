package route

import (
	"fmt"
	"hello/handler"

	"github.com/google/wire"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Routes struct {
	Handler      handler.Handler
	HelloHandler handler.HelloHandler
	AuthHandler  handler.AuthHandler
}

// Handler
func (route Routes) InitRoutes(e *echo.Echo) {
	e.GET("/", route.Handler.Hello)
	e.GET("/hello", route.HelloHandler.Hello)
	e.POST("/login", route.AuthHandler.Login)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/home", handler.Home)

	fmt.Println("Initialize Routes")
}

func NewRoutes(h handler.Handler, helloH handler.HelloHandler, authHandler handler.AuthHandler) Routes {
	return Routes{
		Handler:      h,
		HelloHandler: helloH,
		AuthHandler:  authHandler,
	}
}

var AllHandlers = wire.NewSet(NewRoutes, handler.NewHandler, handler.NewHelloHandler, handler.NewAuthHandler)
