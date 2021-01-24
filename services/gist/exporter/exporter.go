package exporter

import (
	"fmt"
	"github.com/muhammednagy/pipedirve-challenge/db"
	"github.com/muhammednagy/pipedirve-challenge/model"
	"github.com/muhammednagy/pipedirve-challenge/services/github"
	"github.com/muhammednagy/pipedirve-challenge/services/pipedrive"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

// ExportGists will get all persons from DB then get their gists then add new activities to pipedrive
func ExportGists(dbConnection *gorm.DB, config model.Config) {
	people := db.GetPeople(dbConnection, "")
	for personIndex := range people {
		gists, err := github.GetGists(config, people[personIndex].LastVisit, people[personIndex].GithubUsername)
		if err != nil {
			log.Errorf("error while getting user %s gists: %s", people[personIndex].GithubUsername, err)
		} else {
			dbConnection.Model(&people[personIndex]).Update("last_visit", time.Now().UTC())
			for _, gist := range gists {
				var files []model.GistFile
				var notes string
				var subject string
				if gist.GetDescription() != "" {
					notes += fmt.Sprintf("Gist description is %s <br> ",
						gist.GetDescription(),
					)
					subject = gist.GetDescription()
				} else {
					subject = "Github Gist"
				}
				notes += fmt.Sprintf("Gist Pull URL is %s <br> ",
					gist.GetGitPullURL(),
				)
				for _, file := range gist.Files {
					files = append(files, model.GistFile{
						Name:   file.GetFilename(),
						RawURL: file.GetRawURL(),
					})
					notes += fmt.Sprintf("File Name: %s<br> File URL %s<br>", file.GetFilename(), file.GetRawURL())
				}

				dbConnection.Create(&model.Gist{
					Description: gist.GetDescription(),
					PullURL:     gist.GetGitPullURL(),
					PersonID:    people[personIndex].ID,
					Files:       files,
				})
				err = pipedrive.CreateActivity(subject, notes, config.PipedriveToken, people[personIndex].PipedriveID)
				if err != nil {
					log.Errorf("error saving gist activity to pipedrive: %s", err)
				}
			}
		}
	}
}
