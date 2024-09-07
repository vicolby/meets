package types

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name         string `validate:"required"`
	Description  string
	Date         time.Time `validate:"required"`
	LocationID   int       `validate:"required"`
	Location     Location  `validate:"required"`
	OrganizerID  int       `validate:"required"`
	Organizer    User      `validate:"required"`
	Participants []User    `gorm:"many2many:participants;"`
}

type Location struct {
	gorm.Model
	Address   string  `validate:"required"`
	Latitude  float32 `validate:"required"`
	Longitude float32 `validate:"required"`
	City      string  `validate:"required"`
	Country   string  `validate:"required"`
}

type User struct {
	gorm.Model
	FirstName  string `validate:"required"`
	SecondName string
	Email      string
	Phone      string `validate:"required"`
	Rating     int
	RegDate    time.Time `validate:"required"`
}

type AddParticipantReq struct {
	UsersID []int64 `validate:"required"`
}

type DeleteParticipantReq struct {
	UserID int64 `validate:"required"`
}
