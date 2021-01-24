package model

import "time"

type (
	Config struct {
		DBName         string
		GithubToken    string
		PipedriveToken string
		Port           string
	}

	DBModel struct {
		ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
		CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
		UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	}

	Person struct {
		DBModel
		GithubUsername string     `gorm:"size:39;not null;unique;index" json:"github_username"` // Github max length is 39
		LastVisit      *time.Time `json:"last_visit"`
		PipedriveID    uint       `json:"pipedrive_id"`
		Gists          []Gist     `json:"gists" gorm:"OnDelete:SET NULL"`
	}

	Gist struct {
		DBModel
		Description string     `json:"description"`
		PullURL     string     `json:"pull_url"`
		PersonID    uint       `json:"-"`
		Files       []GistFile `json:"files"`
	}

	GistFile struct {
		DBModel
		GistID uint   `json:"-"`
		Name   string `json:"name"`
		RawURL string `json:"raw_url"`
	}
)
