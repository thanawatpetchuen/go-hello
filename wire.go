//+build wireinject

package main

import (
	"hello/app"
	"hello/config"
	"hello/route"
	"hello/server"
	"hello/service"

	"github.com/google/wire"
)

func InitializeApp() (app.App, error) {
	wire.Build(app.NewApp, server.NewServer, config.NewConfig, route.AllHandlers, service.NewService)
	return app.App{}, nil
}
