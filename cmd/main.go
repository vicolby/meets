package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vicolby/meets/data"
	"github.com/vicolby/meets/db"
	"github.com/vicolby/meets/handlers"
)

func main() {
	// Echo instance
	e := echo.New()

	db.InitDB()
	migrateDatabase()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", handlers.CreateUserHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

func migrateDatabase() {
	db.DB.AutoMigrate(&data.User{}, &data.Event{})
	fmt.Println("Database migrations completed.")
}
