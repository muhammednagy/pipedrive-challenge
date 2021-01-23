package main

import (
	"github.com/muhammednagy/pipedirve-challenge/config"
	"github.com/muhammednagy/pipedirve-challenge/db"
	_ "github.com/muhammednagy/pipedirve-challenge/docs"
	"github.com/muhammednagy/pipedirve-challenge/router"
	"github.com/muhammednagy/pipedirve-challenge/services/gist/exporter"
	log "github.com/sirupsen/logrus"
	"time"
)

// @title Pipedrive DevOps Challenge
// @description API to query users gists then save it to
// @contact.name Nagy Salem
// @contact.email me@muhnagy.com
func main() {
	configuration := config.FlagParse()
	dbConnection := db.New(configuration)
	r := router.New(configuration, dbConnection)

	autoFetchGistsTicker := time.NewTicker(3 * time.Hour) // tick once every 3 hours
	// Watch for ticks and with every tick trigger exporting new gists for current users to pipedrive and DB
	go func(ticker *time.Ticker) {
		for ; true; <-ticker.C {
			exporter.ExportGists(dbConnection, configuration)
		}
	}(autoFetchGistsTicker)

	log.Fatal(r.Start("127.0.0.1" + configuration.Port))
}
