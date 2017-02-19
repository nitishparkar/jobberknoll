// CREATE TABLE people (
//    id serial PRIMARY KEY NOT NULL,
//    name varchar(255),
//    bio text
// );

package models

import (
	"errors"
)

type Person struct {
	Id   int
	Name string
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