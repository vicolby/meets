package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func handleBadRequest(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": message})
}

func handleInternalServerError(c echo.Context, message string) error {
	return c.JSON(http.StatusInternalServerError, map[string]string{"error": message})
}

func handleNotFound(c echo.Context, message string) error {
	return c.JSON(http.StatusNotFound, map[string]string{"error": message})
}
