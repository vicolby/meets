package database

import (
	"github.com/vicolby/events/types"
)

func GetUsers() ([]types.User, error) {
	var users []types.User
	if err := DB.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
