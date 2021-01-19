package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite database driver
	"github.com/muhammednagy/pipedirve-challenge/models"
	log "github.com/sirupsen/logrus"
	"os"
)

func New(config models.Config) *gorm.DB {
	DBConnection, err := gorm.Open("sqlite3", config.DBName)
	if err != nil {
		fmt.Println("Cannot connect to  database")
		log.Fatal("storage err: ", err)
	}
	DBConnection.Exec("PRAGMA foreign_keys = ON")
	return DBConnection
}

func TestDB() *gorm.DB {
	DBConnection, err := gorm.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("storage err: ", err)
	}
	DBConnection.Exec("PRAGMA foreign_keys = ON")
	return DBConnection
}

func DropTestDB() error {
	if err := os.Remove("./test.db"); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Person{},
		&models.Gist{},
	)
}
