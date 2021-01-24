package handlers

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/muhammednagy/pipedirve-challenge/db"
	"github.com/muhammednagy/pipedirve-challenge/models"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
	e = NewTestRouter(h)
	_ = loadFixtures()
}

func tearDown() {
	if err := db.DropTestDB(); err != nil {
		log.Fatal(err)
	}
}

func MockPipedrive() {
	httpmock.Activate()
	// Mock creating activity
	httpmock.RegisterResponder("POST", "https://api.pipedrive.com/v1/activities",
		httpmock.NewStringResponder(200,
			`{ "success": true, "data": { "id": 3, "company_id": 7815951, "user_id": 11947605, "done": false, "type": "call", "reference_type": null, "reference_id": null, "conference_meeting_client": null, "conference_meeting_url": null, "due_date": "2021-01-21", "due_time": "", "duration": "", "busy_flag": null, "add_time": "2021-01-21 20:26:29", "marked_as_done_time": "", "last_notification_time": null, "last_notification_user_id": null, "notification_language_id": null, "subject": "test activity", "public_description": null, "calendar_sync_include_context": null, "location": null, "org_id": null, "person_id": 3, "deal_id": null, "lead_id": null, "lead_title": "", "active_flag": true, "update_time": "2021-01-21 20:26:29", "update_user_id": null, "gcal_event_id": null, "google_calendar_id": null, "google_calendar_etag": null, "source_timezone": null, "rec_rule": null, "rec_rule_extension": null, "rec_master_activity_id": null, "conference_meeting_id": null, "note": "test note", "created_by_user_id": 11947605, "location_subpremise": null, "location_street_number": null, "location_route": null, "location_sublocality": null, "location_locality": null, "location_admin_area_level_1": null, "location_admin_area_level_2": null, "location_country": null, "location_postal_code": null, "location_formatted_address": null, "attendees": null, "participants": [ { "person_id": 3, "primary_flag": true } ], "series": null, "org_name": null, "person_name": "test", "deal_title": null, "owner_name": "Nagy", "person_dropbox_bcc": "pipedrivetest-sandbox2@pipedrivemail.com", "deal_dropbox_bcc": null, "assigned_to_user_id": 11947605, "type_name": "Call", "file": null }, "additional_data": { "updates_story_id": 8 }, "related_objects": { "person": { "3": { "active_flag": true, "id": 3, "name": "test", "email": [ { "value": "", "primary": true } ], "phone": [ { "value": "", "primary": true } ] } }, "user": { "11947605": { "id": 11947605, "name": "Nagy", "email": "me@muhnagy.com", "has_pic": 0, "pic_hash": null, "active_flag": true } } } }`,
		))
	// Mock creating activity
	httpmock.RegisterResponder("POST", "https://api.pipedrive.com/v1/persons",
		httpmock.NewStringResponder(200,
			`{ "success": true, "data": { "id": 9, "company_id": 7815951, "owner_id": { "id": 11947605, "name": "Nagy", "email": "me@muhnagy.com", "has_pic": 0, "pic_hash": null, "active_flag": true, "value": 11947605 }, "org_id": null, "name": "muhammednagy3", "first_name": "muhammednagy3", "last_name": null, "open_deals_count": 0, "related_open_deals_count": 0, "closed_deals_count": 0, "related_closed_deals_count": 0, "participant_open_deals_count": 0, "participant_closed_deals_count": 0, "email_messages_count": 0, "activities_count": 0, "done_activities_count": 0, "undone_activities_count": 0, "files_count": 0, "notes_count": 0, "followers_count": 0, "won_deals_count": 0, "related_won_deals_count": 0, "lost_deals_count": 0, "related_lost_deals_count": 0, "active_flag": true, "phone": [ { "value": "", "primary": true } ], "email": [ { "value": "", "primary": true } ], "first_char": "m", "update_time": "2021-01-23 20:27:36", "add_time": "2021-01-23 20:27:36", "visible_to": "3", "picture_id": null, "next_activity_date": null, "next_activity_time": null, "next_activity_id": null, "last_activity_id": null, "last_activity_date": null, "last_incoming_mail_time": null, "last_outgoing_mail_time": null, "label": null, "org_name": null, "cc_email": "pipedrivetest-sandbox2@pipedrivemail.com", "owner_name": "Nagy" }, "related_objects": { "user": { "11947605": { "id": 11947605, "name": "Nagy", "email": "me@muhnagy.com", "has_pic": 0, "pic_hash": null, "active_flag": true } } } }`))
}

func loadFixtures() error {
	p1 := models.Person{
		GithubUsername: "muhammednagy",
	}
	currentTime := time.Now().UTC()
	oneHourBeforeCurrentTime := currentTime.Add(time.Hour * time.Duration(-1))
	oneHourAfterCurrentTime := currentTime.Add(time.Hour * time.Duration(1))
	p2 := models.Person{
		GithubUsername: "muhammednagy2",
		LastVisit:      &currentTime,
		Gists: []models.Gist{{
			DBModel:     models.DBModel{CreatedAt: currentTime},
			Description: "test gist",
			PullURL:     "http://test.dummy/git1",
			Files: []models.GistFile{{
				Name:   "test.go",
				RawURL: "http://test.dummy/test.go",
			}},
		},
			{
				DBModel:     models.DBModel{CreatedAt: oneHourBeforeCurrentTime},
				Description: "test gist made one hour earlier",
				PullURL:     "http://test.dummy/git1",
				Files: []models.GistFile{{
					Name:   "test.go",
					RawURL: "http://test.dummy/test.go",
				}},
			},
			{
				DBModel:     models.DBModel{CreatedAt: oneHourAfterCurrentTime},
				Description: "test gist made one hour later",
				PullURL:     "http://test.dummy/git1",
				Files: []models.GistFile{{
					Name:   "test.go",
					RawURL: "http://test.dummy/test.go",
				}},
			},
		},
	}

	if err := d.Create(&p1).Error; err != nil {
		return err
	}

	if err := d.Create(&p2).Error; err != nil {
		return err
	}

	return nil
}

func TestGetPeople(t *testing.T) {
	tearDown()
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
	tearDown()
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
			MockPipedrive()
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
	tearDown()
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
	tearDown()
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
func NewTestRouter(personHandler *PersonHandler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())

	e.GET("/documentation/*", echoSwagger.WrapHandler)
	apiV1 := e.Group("/api/v1")
	apiV1.GET("/people", personHandler.GetAllPeople)
	apiV1.GET("/person/:username", personHandler.GetPerson)
	apiV1.DELETE("/person/:username", personHandler.DeletePerson)
	apiV1.POST("/person", personHandler.SavePerson)
	return e
}
