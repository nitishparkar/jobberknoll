// CREATE TABLE people (
//    id serial PRIMARY KEY NOT NULL,
//    name varchar(255),
//    bio text
// );

package models

import ()

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
