package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vicolby/events/database"
	"github.com/vicolby/events/types"
)

func TestGetEvents(t *testing.T) {
	cleanupDatabase(t)
	defer cleanupDatabase(t)

	location := types.Location{
		Address:   "Gogol 80",
		Latitude:  15.3,
		Longitude: 76.1,
		City:      "Almaty",
		Country:   "Kazakhstan",
	}
	user := types.User{
		FirstName:  "Alice",
		SecondName: "Johnson",
		Email:      "AJt@mail.com",
		Phone:      "12345",
		Rating:     5,
		RegDate:    time.Now(),
	}

	ctx := context.Background()
	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&location).Error; err != nil {
		tx.Rollback()
		t.Fatalf("Failed to seed database with location: %v", err)
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		t.Fatalf("Failed to seed database with user: %v", err)
	}

	event := types.Event{
		Name:        "BJJ",
		Description: "test",
		Date:        time.Now(),
		LocationID:  int(location.ID),
		OrganizerID: int(user.ID),
	}

	if err := tx.Create(&event).Error; err != nil {
		tx.Rollback()
		t.Fatalf("Failed to seed database with event: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	retrievedEvents, err := database.GetEvents(ctx)

	assert.NoError(t, err, "expected no error from GetEvents")
	assert.Len(t, retrievedEvents, 1, "expected to retrieve 1 event")
	assert.Equal(t, "BJJ", retrievedEvents[0].Name, "expected event name to be 'BJJ'")
}
