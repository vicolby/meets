package database

func Insert[T any](item *T) error {
	result := DB.Create(&item)
	return result.Error
}

func Delete[T any](item *T) error {
	result := DB.Delete(&item)
	return result.Error
}
