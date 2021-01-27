package db

import (
	"fmt"
	"github.com/muhammednagy/pipedrive-challenge/config"
	"github.com/muhammednagy/pipedrive-challenge/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(config config.Config) *gorm.DB {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	dbConnection, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to MySQL database")
		log.Fatal("database connection error: ", err)
	}
	err = AutoMigrate(dbConnection)
	if err != nil {
		fmt.Println("failed to auto migrate error: ", err)
	}
	return dbConnection
}

func TestDB(config config.Config) *gorm.DB {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.TestDBName)
	dbConnection, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection error: ", err)
	}
	return dbConnection
}

func DropTestDB(dbConnection *gorm.DB) error {
	if err := dbConnection.Migrator().DropTable(&model.GistFile{}); err != nil {
		return err
	}
	if err := dbConnection.Migrator().DropTable(&model.Gist{}); err != nil {
		return err
	}
	return dbConnection.Migrator().DropTable(&model.Person{})
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Person{},
		&model.Gist{},
		&model.GistFile{},
	)
}
