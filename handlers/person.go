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

// Get all people
// @Produce  json
// @Summary gets all people who their gists are being monitored
// @Description gets all people who their gists are being monitored
// @Tags Person
// @Success 200 {object} []models.Person
// @Router /api/v1/people [get]
func (h PersonHandler) GetAllPeople(c echo.Context) error {
	people := db.GetPeople(h.db, "")
	return c.JSON(http.StatusOK, people)
}

// Get a specific person
// @Summary gets a specific person based on their username
// @Description gets a specific person based on their username
// @Tags Person
// @Produce  json
// @Param username path string true "github username of the user you want"
// @Success 200 {object} models.Person
// @Failure 404 {object} string
// @Router /api/v1/person/{username} [get]
func (h PersonHandler) GetPerson(c echo.Context) error {
	people := db.GetPeople(h.db, c.Param("username"))
	if len(people) == 0 {
		return c.String(http.StatusNotFound, "Person not found")
	}
	return c.JSON(http.StatusOK, people[0])
}

// Create person
// @Summary Creates person
// @Description Creates person using json
// @Tags Person
// @Accept  json
// @Param models.Person body models.Person  true "assignment Request"
// @Success 201
// @Failure 400 {string} string	"error"
// @Router /api/v1/person [post]
func (h PersonHandler) SavePerson(c echo.Context) error {
	var person models.Person
	if err := c.Bind(&person); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprint("err parsing request: ", err))
	}
	if person.GithubUsername == "" {
		return c.String(http.StatusBadRequest, "missing github username")
	}
	if err := db.SavePerson(h.db, person); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprint("err saving person: ", err))
	}

	return c.NoContent(http.StatusCreated)
}

// Delete person
// @Summary Deletes person
// @Description Deletes person using username
// @Tags Person
// @Param username path string true "github username of the user you want to delete"
// @Success 200
// @Failure 400 {string} string	"error"
// @Router /api/v1/person/{username} [delete]
func (h PersonHandler) DeletePerson(c echo.Context) error {
	err := db.DeletePerson(h.db, c.Param("username"))
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprint("err deleting person: ", err))
	}
	return c.NoContent(http.StatusOK)
}
