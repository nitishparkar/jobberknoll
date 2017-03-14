package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"database/sql/driver"
	"errors"
)

type interactionType string

const (
	CallInteraction interactionType = "call"
	TextInteraction interactionType = "text"
	MeetInteraction interactionType = "meet"
)

// Was getting this error,
// sql: Scan error on column index 5: unsupported Scan, storing driver.Value type []uint8 into type *models.interactionType
// Found solution here - https://github.com/jinzhu/gorm/issues/302
// TODO: Understand this better
func (self *interactionType) Scan(value interface{}) error { *self = interactionType(value.([]byte)); return nil }
func (self interactionType) Value() (driver.Value, error)  { return string(self), nil }

type Interaction struct {
	gorm.Model

	Date time.Time `gorm:"not null"`
	Type interactionType `gorm:"type:varchar(20);not null"`
	Details string `gorm:"not null"`
	PersonID uint `gorm:"not null"`

	Person Person
}

func (self *Interaction) FormattedDate() string {
	// Time formatting is so nice in Go :)
	// But ordinals (2nd Jan) not supported :/
	return self.Date.Format("02 Jan, 2006")
}

func (self *Interaction) FormDate() string {
	return self.Date.Format("02/01/2006")
}

// Scopes
func OrderInteractionDateDesc(db *gorm.DB) *gorm.DB {
	return db.Order("date desc")
}


func SaveInteraction(personId uint, rawIntDate string, intType string, intDetails string) (Interaction, error) {
	db := GetDbConnection()

	intDate, err := time.Parse("02/01/2006", rawIntDate)

	interaction := Interaction{
		Date: intDate,
		Type: interactionType(intType),
		Details: intDetails,
		PersonID: personId,
	}

	if err != nil {
		return interaction, errors.New("Unable to parse interaction date")
	}

	db.Create(&interaction)

	if db.NewRecord(interaction) {
		return interaction, errors.New("Unable to save interaction")
	}

	return interaction, nil
}

func DeleteInteraction(id int) (Interaction, error) {
	db := GetDbConnection()

	interaction := Interaction{}
	db.Delete(&interaction, id)

	return interaction, nil
}

