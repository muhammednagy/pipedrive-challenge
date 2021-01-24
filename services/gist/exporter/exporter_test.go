package exporter

import (
	"github.com/jarcoal/httpmock"
	"github.com/muhammednagy/pipedirve-challenge/db"
	"github.com/muhammednagy/pipedirve-challenge/model"
	"github.com/muhammednagy/pipedirve-challenge/testing/util"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func setup() *gorm.DB {
	d := db.TestDB()
	_ = db.AutoMigrate(d)
	d.Create(&model.Person{
		GithubUsername: "muhammednagy",
		PipedriveID:    6,
	})
	util.MockGithub()
	util.MockPipedrive()
	return d
}

func TestExportGists(t *testing.T) {
	util.TearDown()
	dbConnection := setup()
	defer httpmock.DeactivateAndReset()
	ExportGists(dbConnection, model.Config{})
	p1 := db.GetPeople(dbConnection, "muhammednagy")[0]
	assert.Equal(t, 11, len(p1.Gists))
}
