package main

import (
	"github.com/muhammednagy/pipedirve-challenge/config"
	"github.com/muhammednagy/pipedirve-challenge/router"
)

// @title Pipedrive DevOps Challenge
// @version 0.1.0
// @description API to query users gists then save it to
// @contact.email me@muhnagy.com

func main() {
	configuration := config.FlagParse()

	r := router.New(configuration)
	r.Logger.Fatal(r.Start("127.0.0.1" + configuration.Port))
}
