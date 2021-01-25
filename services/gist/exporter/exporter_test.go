package exporter

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/muhammednagy/pipedrive-challenge/config"
	"github.com/muhammednagy/pipedrive-challenge/db"
	"github.com/muhammednagy/pipedrive-challenge/model"
	"github.com/muhammednagy/pipedrive-challenge/testing/util"
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
	ExportGists(context.Background(), dbConnection, config.Config{})
	p := db.GetPeople(dbConnection, "muhammednagy")
	assert.Equal(t, 1, len(p))
	assert.Equal(t, 11, len(p[0].Gists))
}
