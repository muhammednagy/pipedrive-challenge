package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/muhammednagy/pipedirve-challenge/db"
	"github.com/muhammednagy/pipedirve-challenge/models"
	"net/http"
)

type PersonHandler struct {
	config models.Config
	db     *gorm.DB
}

func NewPersonHandler(config models.Config, db *gorm.DB) *PersonHandler {
	return &PersonHandler{config: config, db: db}
}

func (h PersonHandler) GetAllPeople(c echo.Context) error {
	people := db.GetPeople(h.db, "")
	return c.JSON(http.StatusOK, people)
}

func (h PersonHandler) GetPerson(c echo.Context) error {
	people := db.GetPeople(h.db, c.Param("username"))
	if len(people) == 0 {
		return c.String(http.StatusNotFound, "Person not found")
	}
	return c.JSON(http.StatusOK, people[0])
}
