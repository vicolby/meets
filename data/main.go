package data

import (
	"fmt"
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
	OwnerID      int32
	Participants []int32 `gorm:"type:int[]"`
}

func CreateUser(newUser *User) error {

	db.InitDB()
	newUser.CreatedAt = time.Now()

	result := db.DB.Create(newUser)

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil
}
