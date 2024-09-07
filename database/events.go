package database

import (
	"fmt"

	"github.com/vicolby/events/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetEvents() ([]types.Event, error) {
	var events []types.Event
	if err := DB.Preload("Location").Preload("Organizer").Find(&events).Error; err != nil {
		return events, err
	}

	return events, nil
}

func UpdateEvent(event *types.Event) error {
	if event == nil {
		return fmt.Errorf("event cannot be nil")
	}

	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&types.Event{}, event.ID).Error; err != nil {
			return fmt.Errorf("failed to find event with ID %d: %w", event.ID, err)
		}

		if err := tx.Save(event).Error; err != nil {
			return fmt.Errorf("failed to save event with ID %d: %w", event.ID, err)
		}

		return nil
	})
}

// func AddEventParticipant(eventID int, users []int64) error {
// 	return db.Transaction(func(tx *gorm.DB) error {
// 		var data types.Event

// 		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&data, eventID).Error; err != nil {
// 			return fmt.Errorf("failed to lock event with ID %d: %w", eventID, err)
// 		}

// 		participantSet := make(map[int64]struct{}, len(data.Participants))

// 		for _, participant := range data.Participants {
// 			participantSet[participant] = struct{}{}
// 		}

// 		for _, newParticipant := range users {
// 			if _, ok := participantSet[newParticipant]; ok {
// 				return fmt.Errorf("User %d is already a participant", newParticipant)
// 			}
// 			data.Participants = append(data.Participants, newParticipant)
// 		}

// 		if err := tx.Save(&data).Error; err != nil {
// 			return fmt.Errorf("failed to save participants for event %d: %w", eventID, err)
// 		}

// 		return nil
// 	})
// }

// func DeleteEventParticipant(eventID int, user int64) error {
// 	return db.Transaction(func(tx *gorm.DB) error {
// 		var data types.Event

// 		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&data, eventID).Error; err != nil {
// 			return fmt.Errorf("failed to lock event with ID %d: %w", eventID, err)
// 		}

// 		participantSet := make(map[int64]struct{}, len(data.Participants))

// 		for _, participant := range data.Participants {
// 			participantSet[participant] = struct{}{}
// 		}

// 		if _, ok := participantSet[user]; !ok {
// 			return fmt.Errorf("User %d is not a participant", user)
// 		}

// 		data.Participants = removeValue(data.Participants, user)

// 		if err := tx.Save(&data).Error; err != nil {
// 			return fmt.Errorf("failed to save participants for event %d: %w", eventID, err)
// 		}

// 		return nil
// 	})
// }
