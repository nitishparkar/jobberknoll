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

func (this *Person) setId(value int) {
	this.id = value
}

func (this *Person) SetName(value string) {
	this.name = value
}

func (this *Person) SetBio(value string) {
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

func FetchPerson(id int) (Person, error) {
	db, err := getDbConnection()

	if err == nil {
		defer db.Close()

		person := Person{}
		row := db.QueryRow("SELECT * FROM people WHERE id = $1", id)
		err := row.Scan(&person.id, &person.name, &person.bio)

		if err == nil {
			return person, nil
		} else {
			return person, errors.New("Unable to find person")
		}
	} else {
		return Person{}, errors.New("Unable to get database connection")
	}
}


func SavePerson(name string, bio string) (Person, error) {
	db, err := getDbConnection()

	if err == nil {
		defer db.Close()

		person := Person{}
		person.SetName(name)
		person.SetBio(bio)

		var id int
		err := db.QueryRow("INSERT INTO people(name, bio) VALUES($1, $2) RETURNING id", person.Name(), person.Bio()).Scan(&id)

		if err == nil {
			person.setId(id)
			return person, nil
		} else {
			return person, errors.New("Unable to find person")
		}
	} else {
		return Person{}, errors.New("Unable to get database connection")
	}
}

func UpdatePerson(person Person, name string, bio string) (Person, error) {
	db, err := getDbConnection()

	if err == nil {
		defer db.Close()

		person.SetName(name)
		person.SetBio(bio)

		_, err := db.Exec("UPDATE people SET name= $2, bio= $3 WHERE ID = $1", person.Id(), name, bio)

		if err == nil {
			return person, nil
		} else {
			return person, errors.New("Unable to update record")
		}
	} else {
		return person, errors.New("Unable to get database connection")
	}
}