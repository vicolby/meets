package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/vicolby/meets/data"
	"net/http"
)

func CreateUserHandler(c echo.Context) error {
	// Parse request data
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
