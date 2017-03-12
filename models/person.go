package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model

	Name string `gorm:"not null"`
	Bio  string

	Interactions []Interaction
}

func FetchPeople() ([]Person, error) {
	db := GetDbConnection()

	people := []Person{}
	db.Find(&people)

	return people, nil
}

func FetchPerson(id int) (Person, error) {
	db := GetDbConnection()

	person := Person{}
	db.Preload("Interactions", OrderInteractionDateDesc).First(&person, id)

	return person, nil
}

func SavePerson(name string, bio string) (Person, error) {
	db := GetDbConnection()

	person := Person{Name: name, Bio: bio}
	db.Create(&person)

	if db.NewRecord(person) {
		return person, errors.New("Unable to create person record")
	} else {
		return person, nil
	}
}

func UpdatePerson(person Person, name string, bio string) (Person, error) {
	db := GetDbConnection()

	person.Name = name
	person.Bio = bio

	db.Save(&person)

	return person, nil
}
