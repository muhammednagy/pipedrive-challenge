package db

import (
	"fmt"
	"github.com/muhammednagy/pipedirve-challenge/model"
	"gorm.io/gorm"
)

// GetPeople Gets all persons if not supplied with username
// if supplied with username will return the person with matching username
func GetPeople(dbConnection *gorm.DB, username string) []model.Person {
	var persons []model.Person
	query := dbConnection.Preload("Gists.Files").Table("people")
	if username != "" {
		query = query.Where("github_username = ?", username)
	}
	query.Find(&persons)
	return persons
}

func SavePerson(dbConnection *gorm.DB, person model.Person) error {
	err := dbConnection.Create(&person).Error
	return err
}

func DeletePerson(dbConnection *gorm.DB, username string) error {
	people := GetPeople(dbConnection, username)
	if len(people) == 0 {
		return fmt.Errorf("person not found")
	}
	person := people[0]
	err := dbConnection.Delete(&person).Error
	return err
}
