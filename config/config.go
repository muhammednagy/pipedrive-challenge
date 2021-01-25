package config

import (
	"flag"
	"github.com/muhammednagy/pipedirve-challenge/model"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	buildTime      string
	version        string
	showVersion    = flag.Bool("version", false, "Print version")
	pipedriveToken = flag.String("pipedrive_token", os.Getenv("PIPEDRIVE_TOKEN"), "Pipedrive token")
	GithubToken    = flag.String("github_token", os.Getenv("GITHUB_TOKEN"), "github API token")
	DBName         = flag.String("database_name", os.Getenv("DATABASE_NAME"), "MySQL DB name")
	TestDBName     = flag.String("test_database_name", os.Getenv("TEST_DATABASE_NAME"), "Test DB name")
	DBUsername     = flag.String("database_username", os.Getenv("DATABASE_USERNAME"), "MySQL DB username")
	DBPassword     = flag.String("database_password", os.Getenv("DATABASE_PASSWORD"), "MySQL DB password")
	DBHost         = flag.String("database_host", os.Getenv("DATABASE_HOST"), "MySQL DB host")
	DBPort         = flag.String("database_port", os.Getenv("DATABASE_PORT"), "MySQL DB port")
)

func ParseFlags() model.Config {
	flag.Parse()
	if *showVersion {
		log.Info("Build:", version, buildTime)
		os.Exit(0)
	}

	if *DBName == "" {
		*DBName = "pipedrive"
	}
	if *TestDBName == "" {
		*TestDBName = "test_pipedrive"
	}
	if *DBUsername == "" {
		*DBUsername = "pipedrive"
	}
	if *DBPassword == "" {
		*DBPassword = "pipedrive"
	}
	if *DBHost == "" {
		*DBHost = "127.0.0.1"
	}

	if *DBPort == "" {
		*DBPort = "3306"
	}

	log.Info("Build: " + version + " " + buildTime)
	return model.Config{
		DBName:         *DBName,
		TestDBName:     *TestDBName,
		DBUsername:     *DBUsername,
		DBPassword:     *DBPassword,
		DBHost:         *DBHost,
		DBPort:         *DBPort,
		GithubToken:    *GithubToken,
		PipedriveToken: *pipedriveToken,
	}
}
