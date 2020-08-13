package server

import (
	"fmt"
	"hello/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Echo   *echo.Echo
	Config config.Config
}

func InitServer() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}

func (s Server) StartServer() {
	s.Echo.Logger.Fatal(s.Echo.Start(fmt.Sprintf(":%d", s.Config.Port)))
}

func NewServer(c config.Config) Server {
	e := InitServer()
	return Server{
		Echo:   e,
		Config: c,
	}
}
