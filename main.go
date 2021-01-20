package main

import (
	"github.com/muhammednagy/pipedirve-challenge/config"
	_ "github.com/muhammednagy/pipedirve-challenge/docs"
	"github.com/muhammednagy/pipedirve-challenge/router"
)

// @title Pipedrive DevOps Challenge
// @description API to query users gists then save it to
// @contact.name Nagy Salem
// @contact.email me@muhnagy.com
func main() {
	configuration := config.FlagParse()

	r := router.New(configuration)
	r.Logger.Fatal(r.Start("127.0.0.1" + configuration.Port))
}
