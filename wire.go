//+build wireinject

package main

import (
	"hello/app"
	"hello/route"
	"hello/server"
	"hello/service"

	"github.com/google/wire"
)

func InitializeApp() app.App {
	wire.Build(app.NewApp, server.NewServer, route.AllHandlers, service.NewService)
	return app.App{}
}
