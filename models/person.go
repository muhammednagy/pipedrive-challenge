package models

import "time"

type (
	Config struct {
		DBName         string
		GithubToken    string
		PipedriveToken string
		Port           string
	}

	DBModel struct {
		ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
		CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
		UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	}

	Person struct {
		DBModel
		GithubUsername string    `gorm:"size:39;not null;unique;index" json:"github_username"` // Github max length is 39
		Email          string    `gorm:"size:100;not null;unique" json:"email"`
		LastVisit      time.Time `json:"last_visit"`
		PipedriveID    uint32    `json:"pipedrive_id"`
		Gists          []Gist    `json:"gists"`
	}

	Gist struct {
		DBModel
		RawURL     string `json:"raw_url"`
		PullURL    string `json:"pull_url"`
		ActivityID uint32 `json:"activity_id"`
	}
)
