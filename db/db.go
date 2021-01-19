package db

import (
	"fmt"
	"github.com/muhammednagy/pipedirve-challenge/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite" // sqlite database driver
	"gorm.io/gorm"
	"os"
)

func New(config models.Config) *gorm.DB {
	dbConnection, err := gorm.Open(sqlite.Open(config.DBName), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to  database")
		log.Fatal("storage err: ", err)
	}
	dbConnection.Exec("PRAGMA foreign_keys = ON")
	err = AutoMigrate(dbConnection)
	if err != nil {
		fmt.Println("failed to auto migrate error: ", err)
	}
	return dbConnection
}

func TestDB() *gorm.DB {
	DBConnection, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Person{},
		&models.Gist{},
	)
}
