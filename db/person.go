package db

import (
	"github.com/jinzhu/gorm"
	"github.com/muhammednagy/pipedirve-challenge/models"
)

// Gets all persons if not supplied with username
// if supplied with username will return the person with matching username
func GetPersons(db *gorm.DB, username string) []models.Person {
	var persons []models.Person
	query := db.Table("people")
	if username != "" {
		query = query.Where("username = ?", username)
	}
	query.Find(&persons)
	return persons
}
