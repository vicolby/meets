package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vicolby/events/database"
	"github.com/vicolby/events/types"
)

func TestGetLocations(t *testing.T) {
	locations := []types.Location{
		{
			Address:   "Gogol 80",
			Latitude:  15.3,
			Longitude: 76.1,
			City:      "Almaty",
			Country:   "Kazakhstan",
		},
		{
			Address:   "Gogol 80",
			Latitude:  17.3,
			Longitude: 71.1,
			City:      "Moscow",
			Country:   "Russia",
		},
	}

	ctx := context.Background()

	if err := database.Insert(ctx, &locations); err != nil {
		t.Fatalf("Failed to seed database with locations: %v", err)
	}

	retrievedLocations, err := database.GetLocations(ctx)

	assert.NoError(t, err, "expected no error from GetLocations")
	assert.Len(t, retrievedLocations, 2, "expected to retrieve 2 locations")
	assert.Equal(t, "Almaty", retrievedLocations[0].City, "expected location city to be 'Almaty'")
	assert.Equal(t, "Moscow", retrievedLocations[1].City, "expected location city to be 'Moscow'")
}
