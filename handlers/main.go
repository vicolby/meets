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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	err := data.CreateUser(&newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, newUser)
}

func GetUserHandler(c echo.Context) error {
	userIDstr := c.Param("id")
	userID, err := strconv.ParseUint(userIDstr, 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	user, err := data.GetUser(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUserHandler(c echo.Context) error {
	userIDstr := c.Param("id")
	userID, err := strconv.ParseUint(userIDstr, 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	err = data.DeleteUser(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}

	return c.JSON(http.StatusOK, "User deleted")
}

func CreateEventHandler(c echo.Context) error {
	var newEvent data.Event
	if err := c.Bind(&newEvent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	user, err := data.GetUser(newEvent.OwnerID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid OwnerID"})
	}

	newEvent.Owner = *user

	err = data.CreateEvent(&newEvent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create event"})
	}

	return c.JSON(http.StatusCreated, newEvent)
}
