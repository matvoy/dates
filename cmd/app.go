package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type App struct {
	db     *gorm.DB
	router *echo.Echo
}

func newApp(db *gorm.DB, router *echo.Echo) *App {
	return &App{
		db:     db,
		router: router,
	}
}

func (a *App) init() {
	a.registerRoutes()
	a.initDB()
}
