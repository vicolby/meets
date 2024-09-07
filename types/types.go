package types

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name         string
	Description  string
	Date         time.Time
	LocationID   int
	Location     Location
	OrganizerID  int
	Organizer    User
	Participants []User `gorm:"many2many:participants;"`
}

type Location struct {
	gorm.Model
	Address   string
	Latitude  float32
	Longitude float32
	City      string
	Country   string
}

type User struct {
	gorm.Model
	FirstName  string
	SecondName string
	Email      string
	Phone      string
	Rating     int
	RegDate    time.Time
}

type AddParticipantReq struct {
	UsersID []int64
}

type DeleteParticipantReq struct {
	UserID int64
}
