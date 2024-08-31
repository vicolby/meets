package db

func Insert[T any](item *T) error {
	result := db.Create(&item)
	return result.Error
}

func Delete[T any](item *T) error {
	result := db.Delete(&item)
	return result.Error
}
