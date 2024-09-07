package database

import (
	"github.com/vicolby/events/types"
)

func GetLocations() ([]types.Location, error) {
	var locations []types.Location
	if err := DB.Find(&locations).Error; err != nil {
		return locations, err
	}

	return locations, nil
}
