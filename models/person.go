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
	id   int
	name string
	bio  string
}

func (this *Person) Id() int {
	return this.id
}

func (this *Person) Name() string {
	return this.name
}

func (this *Person) Bio() string {
	return this.bio
}

func (this *Person) setName(value string) {
	this.name = value
}

func (this *Person) setBio(value string) {
	this.bio = value
}

func FetchPeople() ([]Person, error) {
	db, err := getDbConnection()

	if err == nil {
		defer db.Close()
		rows, err := db.Query("SELECT * FROM people")
		if err == nil {
			defer rows.Close()
			result := []Person{}
			for rows.Next() {
				person := Person{}
				err := rows.Scan(&person.id, &person.name, &person.bio)
				if err == nil {
					result = append(result, person)
				}
			}
			return result, nil
		} else {
			return []Person{}, errors.New("Unable to find people")
		}
	} else {
		return []Person{}, errors.New("Unable to get database connection")
	}
}
