package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/muhammednagy/pipedirve-challenge/db"
	"github.com/muhammednagy/pipedirve-challenge/handlers"
	"github.com/muhammednagy/pipedirve-challenge/models"
)

func New(config models.Config) *echo.Echo {
	dbConnection := db.New(config)
	personHandler := handlers.NewPersonHandler(config, dbConnection)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/people", personHandler.GetAllPeople)
	apiV1.GET("/person/:username", personHandler.GetPerson)
	apiV1.DELETE("/person/:username", personHandler.DeletePerson)
	apiV1.POST("/person", personHandler.SavePerson)
	return e
}
