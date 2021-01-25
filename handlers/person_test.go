package handlers

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo/v4"
	"github.com/muhammednagy/pipedrive-challenge/config"
	"github.com/muhammednagy/pipedrive-challenge/db"
	"github.com/muhammednagy/pipedrive-challenge/model"
	"github.com/muhammednagy/pipedrive-challenge/testing/util"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	dbConnection  *gorm.DB
	personHandler *PersonHandler
	e             *echo.Echo
)

func setup() {
	dbConnection = db.TestDB(config.ParseFlags())
	util.TearDown(dbConnection)
	_ = db.AutoMigrate(dbConnection)
	personHandler = NewPersonHandler(config.Config{}, dbConnection)
	e = echo.New()
	_ = util.LoadFixtures(dbConnection)
}

func TestGetPeople(t *testing.T) {
	setup()
	req := httptest.NewRequest(echo.GET, "/api/v1/people", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, personHandler.GetAllPeople(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var response []model.Person
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(response))
	}
}

func TestSavePerson(t *testing.T) {
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
			util.MockPipedrive()
		}
		req := httptest.NewRequest(echo.GET, "/api/v1/people", nil)
		q := req.URL.Query()
		q.Add("username", tp.username)
		req.URL.RawQuery = q.Encode()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, personHandler.SavePerson(c))
		assert.Equal(t, tp.statusCode, rec.Code)
	}
}

func TestDeletePerson(t *testing.T) {
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
		req := httptest.NewRequest(echo.DELETE, "/api/v1/people/"+tp.username, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("username")
		c.SetParamValues(tp.username)
		assert.NoError(t, personHandler.DeletePerson(c))
		assert.Equal(t, tp.statusCode, rec.Code)
	}
}

func TestGetPerson(t *testing.T) {
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
		req := httptest.NewRequest(echo.GET, "/api/v1/people/"+tp.username, nil)
		q := req.URL.Query()
		for k, v := range tp.params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("username")
		c.SetParamValues(tp.username)
		assert.NoError(t, personHandler.GetPerson(c))
		if assert.Equal(t, tp.statusCode, rec.Code) && tp.statusCode != http.StatusNotFound {
			var response model.Person
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tp.responseLength, len(response.Gists))
		}
	}
}
