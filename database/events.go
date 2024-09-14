package database

import (
	"context"
	"fmt"

	"github.com/vicolby/events/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetEvents(ctx context.Context) ([]types.Event, error) {
	var events []types.Event
	if err := DB.WithContext(ctx).Preload("Location").Preload("Organizer").Find(&events).Error; err != nil {
		return events, err
	}

	return events, nil
}

func UpdateEvent(ctx context.Context, event *types.Event) error {
	if event == nil {
		return fmt.Errorf("event cannot be nil")
	}

	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&types.Event{}, event.ID).Error; err != nil {
			return fmt.Errorf("failed to find event with ID %d: %w", event.ID, err)
		}

		if err := tx.Save(event).Error; err != nil {
			return fmt.Errorf("failed to save event with ID %d: %w", event.ID, err)
		}

		return nil
	})
}
