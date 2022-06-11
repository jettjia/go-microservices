// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"go-micro-frame-doc/25-wire/05-wire-interface/internal/config"
	"go-micro-frame-doc/25-wire/05-wire-interface/internal/db"
)

// Injectors from wire.go:

//go:generate wire
func InitApp() (*App, error) {
	configConfig, err := config.New()
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.NewDb(configConfig)
	if err != nil {
		return nil, err
	}
	dao := db.NewDao(sqlDB)
	app := NewApp(dao)
	return app, nil
}
