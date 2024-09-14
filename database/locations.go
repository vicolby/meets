package database

import (
	"context"

	"github.com/vicolby/events/types"
)

func GetLocations(ctx context.Context) ([]types.Location, error) {
	var locations []types.Location
	if err := DB.WithContext(ctx).Find(&locations).Error; err != nil {
		return locations, err
	}

	return locations, nil
}
