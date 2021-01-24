package exporter

import (
	"github.com/jarcoal/httpmock"
	"github.com/muhammednagy/pipedirve-challenge/db"
	"github.com/muhammednagy/pipedirve-challenge/models"
	"github.com/muhammednagy/pipedirve-challenge/testing/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func setup() *gorm.DB {
	d := db.TestDB()
	_ = db.AutoMigrate(d)
	d.Create(&models.Person{
		GithubUsername: "muhammednagy",
		PipedriveID:    6,
	})
	utils.MockGithub()
	utils.MockPipedrive()
	return d
}

func TestExportGists(t *testing.T) {
	utils.TearDown()
	dbConnection := setup()
	defer httpmock.DeactivateAndReset()
	ExportGists(dbConnection, models.Config{})
	p1 := db.GetPeople(dbConnection, "muhammednagy")[0]
	assert.Equal(t, 11, len(p1.Gists))
}
