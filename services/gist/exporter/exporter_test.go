package exporter

import (
	"github.com/jarcoal/httpmock"
	"github.com/muhammednagy/pipedirve-challenge/config"
	"github.com/muhammednagy/pipedirve-challenge/db"
	"github.com/muhammednagy/pipedirve-challenge/model"
	"github.com/muhammednagy/pipedirve-challenge/testing/util"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func setup() *gorm.DB {
	d := db.TestDB(config.ParseFlags())
	util.TearDown(d)
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
	dbConnection := setup()
	defer httpmock.DeactivateAndReset()
	ExportGists(dbConnection, model.Config{})
	p := db.GetPeople(dbConnection, "muhammednagy")
	assert.Equal(t, 1, len(p))
	assert.Equal(t, 11, len(p[0].Gists))
}
