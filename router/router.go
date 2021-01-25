package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/muhammednagy/pipedrive-challenge/handlers"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func New(personHandler *handlers.PersonHandler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())

	e.GET("/documentation/*", echoSwagger.WrapHandler)
	apiV1 := e.Group("/api/v1")
	apiV1.GET("/people", personHandler.GetAllPeople)
	apiV1.GET("/people/:username", personHandler.GetPerson)
	apiV1.DELETE("/people/:username", personHandler.DeletePerson)
	apiV1.POST("/people", personHandler.SavePerson)
	return e
}
