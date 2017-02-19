package models

import (
	"github.com/jinzhu/gorm"
	"errors"
)

func RunMigrations() {
	db, err := getDbConnection()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	db.AutoMigrate(&Person{})
}

type Person struct {
	gorm.Model

	Name string	`gorm:"not null"`
	Bio  string
}

func FetchPeople() ([]Person, error) {
	db, err := getDbConnection()

	if err == nil {
		defer db.Close()

		people := []Person{}
		db.Find(&people)

		return people, nil
	} else {
		return []Person{}, errors.New("Unable to get database connection")
	}
}

func FetchPerson(id int) (Person, error) {
	db, err := getDbConnection()

	if err == nil {
		defer db.Close()

		person := Person{}
		db.First(&person, id)

		return person, nil
	} else {
		return Person{}, errors.New("Unable to get database connection")
	}
}


func SavePerson(name string, bio string) (Person, error) {
	db, err := getDbConnection()

	if err == nil {
		defer db.Close()

		person := Person{Name: name, Bio: bio}

		db.Create(&person)

		if db.NewRecord(person) {
			return person, errors.New("Unable to create person record")
		} else {
			return person, nil
		}
	} else {
		return Person{}, errors.New("Unable to get database connection")
	}
}

func UpdatePerson(person Person, name string, bio string) (Person, error) {
	db, err := getDbConnection()

	if err == nil {
		defer db.Close()

		person.Name = name
		person.Bio = bio

		db.Save(&person)

		return person, nil
	} else {
		return person, errors.New("Unable to get database connection")
	}
}