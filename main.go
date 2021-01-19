package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/muhammednagy/pipedirve-challenge/config"
	"github.com/muhammednagy/pipedirve-challenge/db"
	log "github.com/sirupsen/logrus"
)

// @title Pipedrive DevOps Challenge
// @version 0.1.0
// @description API to query users gists then save it to
// @contact.email me@muhnagy.com

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	configuration := config.FlagParse()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	db.AutoMigrate(db.New(configuration))
}
