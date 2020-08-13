//+build wireinject

package main

import (
	"hello/app"
	"hello/handler"

	"github.com/google/wire"
)

func InitializeApp() app.App {
	wire.Build(app.NewApp, handler.HandlerSet)
	return app.App{}
}
