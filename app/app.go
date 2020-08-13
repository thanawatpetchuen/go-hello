package app

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"hello/config"
	"hello/route"
)

type App struct {
	AppName string
	Echo    *echo.Echo
}

func (a *App) InitEcho() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	route.InitRoutes(e)

	c, _ := config.GetConfig()
	fmt.Println(c.Port)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", c.Port)))

	a.Echo = e

	return e
}

func NewApp() App {
	return App{}
}
