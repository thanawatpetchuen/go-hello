// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"hello/app"
)

// Injectors from wire.go:

func InitializeApp() app.App {
	appApp := app.NewApp()
	return appApp
}
