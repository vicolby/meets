package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vicolby/events/database"
	"github.com/vicolby/events/types"
)

func TestGetUsers(t *testing.T) {
	cleanupDatabase(t)
	defer cleanupDatabase(t)

	users := []types.User{
		{
			FirstName:  "Alice",
			SecondName: "Johnson",
			Email:      "AJt@mail.com",
			Phone:      "12345",
			Rating:     5,
			RegDate:    time.Now(),
		},
		{
			FirstName:  "John",
			SecondName: "Born",
			Email:      "JBt@mail.com",
			Phone:      "54321",
			Rating:     4,
			RegDate:    time.Now(),
		},
	}

	ctx := context.Background()
	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&users).Error; err != nil {
		tx.Rollback()
		t.Fatalf("failed to seed database: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	retrievedUsers, err := database.GetUsers(ctx)

	assert.NoError(t, err, "expected no error from GetUsers")
	assert.Len(t, retrievedUsers, 2, "expected to retrieve 2 users")
	assert.Equal(t, "Alice", retrievedUsers[0].FirstName, "expected user name to be 'Alice'")
	assert.Equal(t, "John", retrievedUsers[1].FirstName, "expected user name to be 'John'")
}
