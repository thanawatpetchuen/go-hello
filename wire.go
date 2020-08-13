//+build wireinject

package main

import (
	"hello/app"

	"github.com/google/wire"
)

func InitializeApp() app.App {
	wire.Build(app.NewApp)
	return app.App{}
}
