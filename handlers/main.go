package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/vicolby/meets/data"
	"net/http"
	"strconv"
)

func CreateUserHandler(c echo.Context) error {
	var newUser data.User
	if err := c.Bind(&newUser); err != nil {
		return handleBadRequest(c, "Invalid request payload")
	}

	err := data.CreateUser(&newUser)
	if err != nil {
		return handleInternalServerError(c, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, newUser)
}

func GetUserHandler(c echo.Context) error {
	userIDstr := c.Param("id")
	userID, err := strconv.ParseUint(userIDstr, 10, 64)

	if err != nil {
		return handleBadRequest(c, "Invalid user ID")
	}

	user, err := data.GetUser(uint(userID))
	if err != nil {
		return handleInternalServerError(c, "Failed to get user")
	}

	if user == nil {
		return handleNotFound(c, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUserHandler(c echo.Context) error {
	userIDstr := c.Param("id")
	userID, err := strconv.ParseUint(userIDstr, 10, 64)

	if err != nil {
		return handleBadRequest(c, "Invalid user ID")
	}

	err = data.DeleteUser(uint(userID))
	if err != nil {
		return handleInternalServerError(c, "Failed to delete user")
	}

	return c.JSON(http.StatusOK, "User deleted")
}

func CreateEventHandler(c echo.Context) error {
	var newEvent data.Event

	if err := c.Bind(&newEvent); err != nil {
		return handleBadRequest(c, "Invalid event ID")
	}

	user, err := data.GetUser(newEvent.OwnerID)
	if err != nil {
		return handleBadRequest(c, "Invalid OwnerID")
	}

	newEvent.Owner = *user

	err = data.CreateEvent(&newEvent)
	if err != nil {
		return handleInternalServerError(c, "Failed to get user")
	}

	return c.JSON(http.StatusCreated, newEvent)
}

func DeleteEventHandler(c echo.Context) error {
	eventIDstr := c.Param("id")
	eventID, err := strconv.ParseUint(eventIDstr, 10, 64)

	if err != nil {
		return handleBadRequest(c, "Invalid event ID")
	}

	err = data.DeleteEvent(uint(eventID))
	if err != nil {
		return handleInternalServerError(c, "Failed to delete event")
	}

	return c.JSON(http.StatusOK, "Event deleted")
}
