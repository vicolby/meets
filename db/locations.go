package db

import (
	"github.com/vicolby/events/types"
)

func GetLocations() ([]types.Location, error) {
	var locations []types.Location
	if err := db.Find(&locations).Error; err != nil {
		return locations, err
	}

	return locations, nil
}
