package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/muhammednagy/pipedirve-challenge/db"
	"github.com/muhammednagy/pipedirve-challenge/models"
	"gorm.io/gorm"
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

func (h PersonHandler) SavePerson(c echo.Context) error {
	var person models.Person
	if err := c.Bind(&person); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprint("err parsing request: ", err))
	}
	if err := db.SavePerson(h.db, person); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprint("err saving person: ", err))
	}

	return c.NoContent(http.StatusCreated)
}
