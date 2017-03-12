package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type interactionType string

const (
	call interactionType = "call"
	text interactionType = "text"
	meet interactionType = "meet"
)

type Interaction struct {
	gorm.Model

	Date time.Time `gorm:"not null"`
	Type interactionType `gorm:"type:varchar(20);not null"`
	Details string `gorm:"not null"`
	PersonID int `gorm:"not null"`

	Person Person
}