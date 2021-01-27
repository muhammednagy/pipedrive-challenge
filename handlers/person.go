package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/muhammednagy/pipedrive-challenge/config"
	"github.com/muhammednagy/pipedrive-challenge/db"
	"github.com/muhammednagy/pipedrive-challenge/model"
	"github.com/muhammednagy/pipedrive-challenge/services/pipedrive"
	"gorm.io/gorm"
	"net/http"
)

type PersonHandler struct {
	config config.Config
	db     *gorm.DB
}

func NewPersonHandler(config config.Config, db *gorm.DB) *PersonHandler {
	return &PersonHandler{config: config, db: db}
}

// Get all people
// @Produce  json
// @Summary gets all people who their gists are being monitored
// @Description gets all people who their gists are being monitored
// @Tags Person
// @Success 200 {object} []model.Person
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
// @Param getAllGists query bool false "get all gists not only the ones added since last visit"
// @Success 200 {object} model.Person
// @Failure 404 {object} string
// @Router /api/v1/people/{username} [get]
func (h PersonHandler) GetPerson(c echo.Context) error {
	people := db.GetPeople(h.db, c.Param("username"))
	if len(people) == 0 {
		return c.String(http.StatusNotFound, "Person not found")
	}
	var gistsSinceLastVisit []model.Gist
	person := people[0]
	if person.LastVisit != nil && c.QueryParam("getAllGists") != "true" {
		for _, gist := range person.Gists {
			if sinceLastVisit(gist, person) {
				gistsSinceLastVisit = append(gistsSinceLastVisit, gist)
			}
		}
		person.Gists = gistsSinceLastVisit
	}
	return c.JSON(http.StatusOK, person)
}

// Create person
// @Summary Creates person
// @Description Creates person using json
// @Tags Person
// @Accept  x-www-form-urlencoded
// @Param username formData string  true "username"
// @Success 201
// @Failure 400 {string} string	"error"
// @Router /api/v1/people [post]
func (h PersonHandler) SavePerson(c echo.Context) error {
	username := c.FormValue("username")
	if username == "" {
		return c.String(http.StatusBadRequest, "missing github username")
	}
	pipedrivePersonID, err := pipedrive.CreatePerson(username, h.config.PipedriveToken)
	if pipedrivePersonID == 0 {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error creating a person in pipedrive: %s", err))
	}
	person := model.Person{GithubUsername: username, PipedriveID: uint(pipedrivePersonID)}
	if err := db.SavePerson(h.db, person); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprint("error saving person: ", err))
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
// @Router /api/v1/people/{username} [delete]
func (h PersonHandler) DeletePerson(c echo.Context) error {
	err := db.DeletePerson(h.db, c.Param("username"))
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprint("err deleting person: ", err))
	}
	return c.NoContent(http.StatusOK)
}

func sinceLastVisit(gist model.Gist, person model.Person) bool {
	// convert to unix to avoid returning false if there is milliseconds difference
	if gist.CreatedAt.Unix() == person.LastVisit.Unix() || gist.CreatedAt.After(*person.LastVisit) {
		return true
	}
	return false
}
