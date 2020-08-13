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

func (s *Server) InitServer() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	c, _ := config.GetConfig()

	// Start server
	s.Echo = e
	s.Config = c
}

func (s Server) StartServer() {
	s.Echo.Logger.Fatal(s.Echo.Start(fmt.Sprintf(":%d", s.Config.Port)))
}

func NewServer() Server {
	return Server{}
}
