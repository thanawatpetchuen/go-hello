package app

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"hello/config"
	"hello/handler"
	"hello/route"
)

type App struct {
	AppName string
	Echo    *echo.Echo
	Handler handler.Handler
}

func (a *App) InitEcho() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	route.InitRoutes(e, a.Handler)

	c, _ := config.GetConfig()
	fmt.Println(c.Port)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", c.Port)))

	a.Echo = e

	return e
}

func NewApp(h handler.Handler) App {
	return App{
		Handler: h,
	}
}
