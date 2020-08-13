package app

import (
	"hello/route"
	"hello/server"
)

type App struct {
	AppName string
	Server  server.Server
	Routes  route.Routes
}

func (a *App) StartApp() {
	a.Server.InitServer()
	a.Routes.InitRoutes(a.Server.Echo)
	a.Server.StartServer()
}

func NewApp(s server.Server, r route.Routes) App {
	return App{
		Server: s,
		Routes: r,
	}
}
