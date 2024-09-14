package database

import "context"

func Insert[T any](ctx context.Context, item *T) error {
	result := DB.WithContext(ctx).Create(&item)
	return result.Error
}

func Delete[T any](ctx context.Context, item *T) error {
	result := DB.WithContext(ctx).Delete(&item)
	return result.Error
}
