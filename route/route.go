package route

import (
	"fmt"
	"hello/handler"
	"hello/server"

	"github.com/google/wire"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4/middleware"
)

type Routes struct {
	Handler         handler.Handler
	HelloHandler    handler.HelloHandler
	AuthHandler     handler.AuthHandler
	RestrictHandler handler.RestrictHandler
	Server          server.Server
}

// Root Handler
func (route Routes) InitRoutes() {
	route.Server.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(route.Server.Echo)

	fmt.Println("Initialize Routes")
}

func NewRoutes(s server.Server, h handler.Handler, helloH handler.HelloHandler, authHandler handler.AuthHandler, res handler.RestrictHandler) Routes {
	return Routes{
		Handler:         h,
		HelloHandler:    helloH,
		AuthHandler:     authHandler,
		RestrictHandler: res,
		Server:          s,
	}
}

var AllHandlers = wire.NewSet(NewRoutes, handler.NewHandler, handler.NewHelloHandler, handler.NewAuthHandler, handler.NewRestrictHandler)
