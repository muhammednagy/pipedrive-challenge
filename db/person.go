package db

import (
	"github.com/muhammednagy/pipedirve-challenge/models"
	"gorm.io/gorm"
)

// Gets all persons if not supplied with username
// if supplied with username will return the person with matching username
func GetPeople(dbConnection *gorm.DB, username string) []models.Person {
	var persons []models.Person
	query := dbConnection.Table("people")
	if username != "" {
		query = query.Where("github_username = ?", username)
	}
	query.Find(&persons)
	return persons
}

func SavePerson(dbConnection *gorm.DB, person models.Person) error {
	err := dbConnection.Create(&person).Error
	return err
}
