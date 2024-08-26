package db

import "github.com/vicolby/events/types"

func InsertEvent(event *types.Event) error {
	result := db.Create(&event)
	return result.Error
}

func DeleteEvent(event *types.Event) error {
	result := db.Delete(&event)
	return result.Error
}

func GetEvents() ([]types.Event, error) {
	var events []types.Event
	if err := db.Preload("Location").Preload("Organizer").Find(&events).Error; err != nil {
		return events, err
	}

	return events, nil
}

func InsertLocation(location *types.Location) error {
	result := db.Create(&location)
	return result.Error
}

func DeleteLocation(location *types.Location) error {
	result := db.Delete(&location)
	return result.Error
}

func GetLocations() ([]types.Location, error) {
	var locations []types.Location
	if err := db.Find(&locations).Error; err != nil {
		return locations, err
	}

	return locations, nil
}

func InsertUser(user *types.User) error {
	result := db.Create(&user)
	return result.Error
}

func DeleteUser(user *types.User) error {
	result := db.Delete(&user)
	return result.Error
}

func GetUsers() ([]types.User, error) {
	var users []types.User
	if err := db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
