package handlers

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo/v4"
	"github.com/muhammednagy/pipedirve-challenge/db"
	"github.com/muhammednagy/pipedirve-challenge/models"
	"github.com/muhammednagy/pipedirve-challenge/testing/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	d *gorm.DB
	h *PersonHandler
	e *echo.Echo
)

func setup() {
	d = db.TestDB()
	_ = db.AutoMigrate(d)
	h = NewPersonHandler(models.Config{}, d)
	e = echo.New()
	_ = utils.LoadFixtures(d)
}

func TestGetPeople(t *testing.T) {
	utils.TearDown()
	setup()
	req := httptest.NewRequest(echo.GET, "/api/v1/people", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, h.GetAllPeople(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var response []models.Person
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(response))
	}
}

func TestSavePerson(t *testing.T) {
	utils.TearDown()
	setup()
	testCases := map[string]struct {
		statusCode int
		username   string
		mocked     bool
	}{
		"PipedriveFails": {
			http.StatusInternalServerError,
			"muhammednagy3",
			false,
		},
		"NewPerson": {
			http.StatusCreated,
			"muhammednagy4",
			true,
		},
		"AlreadyExists": {
			http.StatusBadRequest,
			"muhammednagy",
			true,
		},
		"NoUsername": {
			http.StatusBadRequest,
			"",
			true,
		},
	}
	defer httpmock.DeactivateAndReset()
	for _, tp := range testCases {
		if tp.mocked {
			utils.MockPipedrive()
		}
		req := httptest.NewRequest(echo.GET, "/api/v1/person", nil)
		q := req.URL.Query()
		q.Add("username", tp.username)
		req.URL.RawQuery = q.Encode()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, h.SavePerson(c))
		assert.Equal(t, tp.statusCode, rec.Code)
	}
}

func TestDeletePerson(t *testing.T) {
	utils.TearDown()
	setup()
	testCases := map[string]struct {
		statusCode int
		username   string
	}{
		"ExistingPerson": {
			http.StatusOK,
			"muhammednagy",
		},
		"doesNotExist": {
			http.StatusBadRequest,
			"NotExistent",
		},
	}

	for _, tp := range testCases {
		req := httptest.NewRequest(echo.DELETE, "/api/v1/person/"+tp.username, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("username")
		c.SetParamValues(tp.username)
		assert.NoError(t, h.DeletePerson(c))
		assert.Equal(t, tp.statusCode, rec.Code)
	}
}

func TestGetPerson(t *testing.T) {
	utils.TearDown()
	setup()
	testCases := map[string]struct {
		params         map[string]string
		statusCode     int
		responseLength int
		username       string
	}{
		"WithGistsSinceLastVisit": {
			map[string]string{},
			http.StatusOK,
			2,
			"muhammednagy2",
		},
		"WithAllGists": {
			map[string]string{
				"getAllGists": "true",
			},
			http.StatusOK,
			3,
			"muhammednagy2",
		},
		"withoutGists": {
			map[string]string{},
			http.StatusOK,
			0,
			"muhammednagy",
		},
		"doesNotExist": {
			map[string]string{},
			http.StatusNotFound,
			0,
			"NotExistent",
		},
	}

	for _, tp := range testCases {
		req := httptest.NewRequest(echo.GET, "/api/v1/person/"+tp.username, nil)
		q := req.URL.Query()
		for k, v := range tp.params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("username")
		c.SetParamValues(tp.username)
		assert.NoError(t, h.GetPerson(c))
		if assert.Equal(t, tp.statusCode, rec.Code) && tp.statusCode != http.StatusNotFound {
			var response models.Person
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tp.responseLength, len(response.Gists))
		}
	}
}
