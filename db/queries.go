package db

import (
	"fmt"

	"github.com/vicolby/events/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ---------------Events------------------------

func InsertEvent(event *types.Event) error {
	result := db.Create(&event)
	return result.Error
}

func DeleteEvent(event *types.Event) error {
	result := db.Delete(&event)
	return result.Error
}

func GetEvents() ([]types.Event, error) {
	var events []types.Event
	if err := db.Preload("Location").Preload("Organizer").Find(&events).Error; err != nil {
		return events, err
	}

	return events, nil
}

func AddEventParticipant(eventID int, users []int64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var data types.Event

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&data, eventID).Error; err != nil {
			return fmt.Errorf("failed to lock event with ID %d: %w", eventID, err)
		}

		participantSet := make(map[int64]struct{}, len(data.Participants))

		for _, participant := range data.Participants {
			participantSet[participant] = struct{}{}
		}

		for _, newParticipant := range users {
			if _, ok := participantSet[newParticipant]; ok {
				return fmt.Errorf("User %d is already a participant", newParticipant)
			}
			data.Participants = append(data.Participants, newParticipant)
		}

		if err := tx.Save(&data).Error; err != nil {
			return fmt.Errorf("failed to save participants for event %d: %w", eventID, err)
		}

		return nil
	})
}

func DeleteEventParticipant(eventID int, user int64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var data types.Event

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&data, eventID).Error; err != nil {
			return fmt.Errorf("failed to lock event with ID %d: %w", eventID, err)
		}

		participantSet := make(map[int64]struct{}, len(data.Participants))

		for _, participant := range data.Participants {
			participantSet[participant] = struct{}{}
		}

		if _, ok := participantSet[user]; !ok {
			return fmt.Errorf("User %d is not a participant", user)
		}

		data.Participants = removeValue(data.Participants, user)

		if err := tx.Save(&data).Error; err != nil {
			return fmt.Errorf("failed to save participants for event %d: %w", eventID, err)
		}

		return nil
	})
}

// -----------------Locations----------------------

func InsertLocation(location *types.Location) error {
	result := db.Create(&location)
	return result.Error
}

func DeleteLocation(location *types.Location) error {
	result := db.Delete(&location)
	return result.Error
}

func GetLocations() ([]types.Location, error) {
	var locations []types.Location
	if err := db.Find(&locations).Error; err != nil {
		return locations, err
	}

	return locations, nil
}

//-----------------------Users------------------------------

func InsertUser(user *types.User) error {
	result := db.Create(&user)
	return result.Error
}

func DeleteUser(user *types.User) error {
	result := db.Delete(&user)
	return result.Error
}

func GetUsers() ([]types.User, error) {
	var users []types.User
	if err := db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
