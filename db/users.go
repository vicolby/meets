package db

import (
	"github.com/vicolby/events/types"
)

func GetUsers() ([]types.User, error) {
	var users []types.User
	if err := db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
