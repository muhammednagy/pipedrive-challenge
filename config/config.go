package config

import (
	"flag"
	"github.com/muhammednagy/pipedirve-challenge/models"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	buildTime      string
	version        string
	showVersion    = flag.Bool("version", false, "Print version")
	pipedriveToken = flag.String("pipedrive token", os.Getenv("PIPEDRIVE_TOKEN"), "Pipedrive token")
	GithubToken    = flag.String("github token", os.Getenv("GITHUB_TOKEN"), "github API token")
	DBName         = flag.String("database name", os.Getenv("DATABASE_NAME"), "Sqlite DB name")
	Port           = flag.String("application listening port", os.Getenv("PORT"), "Listen to port")
)

func FlagParse() models.Config {
	flag.Parse()
	if *showVersion {
		log.Info("Build:", version, buildTime)
		os.Exit(0)
	}

	if *pipedriveToken == "" {
		log.Fatal("Pipedrive Token is required!")
	}
	if *DBName == "" {
		*DBName = "database.sqlite3"
	}

	if *Port == "" {
		*Port = "3000"
	}

	log.Info("Build: " + version + " " + buildTime)
	return models.Config{
		DBName:         *DBName,
		GithubToken:    *GithubToken,
		PipedriveToken: *pipedriveToken,
		Port:           ":" + *Port,
	}
}
