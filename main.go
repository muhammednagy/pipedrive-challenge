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
	getGistsCtx := context.Background() // context to be used with get gists. can be used to set timeouts for github API in the future
	dbConnection := db.New(configuration)
	if configuration.FetchNewGists {
		exporter.ExportGists(getGistsCtx, dbConnection, configuration)
		log.Info("Finished fetching new gists")
	} else {
		personHandler := handlers.NewPersonHandler(configuration, dbConnection)
		r := router.New(personHandler)
		log.Fatal(r.Start("0.0.0.0:3000"))
	}

}
