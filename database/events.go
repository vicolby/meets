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
