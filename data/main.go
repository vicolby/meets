package data

import (
	"fmt"
	"github.com/lib/pq"
	"github.com/vicolby/meets/db"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string
}

type Event struct {
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string
	OwnerID      uint
	Owner        User
	Participants pq.Int64Array `gorm:"type:int[]"`
}

func CreateUser(newUser *User) error {
	newUser.CreatedAt = time.Now()

	result := db.DB.Create(newUser)

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil
}

func GetUser(id uint) (*User, error) {
	var user User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func DeleteUser(id uint) error {
	result := db.DB.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetEvents() ([]Event, error) {
	var events []Event
	result := db.DB.Find(&events)
	if result.Error != nil {
		return nil, result.Error
	}
	return events, nil
}

func CreateEvent(newEvent *Event) error {
	newEvent.CreatedAt = time.Now()

	result := db.DB.Preload("Owner").Create(newEvent)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteEvent(id uint) error {
	result := db.DB.Delete(&Event{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateEventName(eventID uint, newName string) error {
	var existingEvent Event
	if err := db.DB.First(&existingEvent, eventID).Error; err != nil {
		return err
	}
	existingEvent.Name = newName

	if err := db.DB.Save(&existingEvent).Error; err != nil {
		return err
	}

	return nil
}
