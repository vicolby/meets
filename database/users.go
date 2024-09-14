package database

import (
	"context"

	"github.com/vicolby/events/types"
)

func GetUsers(ctx context.Context) ([]types.User, error) {
	var users []types.User
	if err := DB.WithContext(ctx).Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
