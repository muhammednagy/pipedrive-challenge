package main

import (
	"context"
	"github.com/muhammednagy/pipedrive-challenge/config"
	"github.com/muhammednagy/pipedrive-challenge/db"
	_ "github.com/muhammednagy/pipedrive-challenge/docs"
	"github.com/muhammednagy/pipedrive-challenge/handlers"
	"github.com/muhammednagy/pipedrive-challenge/router"
	"github.com/muhammednagy/pipedrive-challenge/services/gist/exporter"
	log "github.com/sirupsen/logrus"
	"time"
)

// @title Pipedrive DevOps Challenge
// @description API to query users gists then save it to
// @contact.name Nagy Salem
// @contact.email me@muhnagy.com
func main() {
	configuration := config.ParseFlags()
	if configuration.PipedriveToken == "" {
		log.Fatal("Pipedrive Token is required!")
	}
	dbConnection := db.New(configuration)
	personHandler := handlers.NewPersonHandler(configuration, dbConnection)
	r := router.New(personHandler)

	autoFetchGistsTicker := time.NewTicker(3 * time.Hour) // tick once every 3 hours
	getGistsCtx := context.Background()                   // context to be used with get gists. can be used to set timeouts for github API in the future
	// Watch for ticks and with every tick trigger exporting new gists for current users to pipedrive and DB
	go func(ticker *time.Ticker) {
		for ; true; <-ticker.C {
			exporter.ExportGists(getGistsCtx, dbConnection, configuration)
		}
	}(autoFetchGistsTicker)

	log.Fatal(r.Start("0.0.0.0:3000"))
}
