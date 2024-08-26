package main

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/vicolby/events/db"
	"github.com/vicolby/events/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	err = db.Init()

	if err != nil {
		slog.Info("Error connecting to database: %v\n", err)
		return
	}

	slog.Info("Successfully connected to the database!")

	r := chi.NewRouter()

	r.Get("/api/v1/events", handlers.GetEventsHandler)
	r.Post("/api/v1/events", handlers.CreateEventHandler)
	r.Delete("/api/v1/events", handlers.DeleteEventHandler)

	r.Group(func(r chi.Router) {
		r.Use(handlers.EnsureValidToken())
		r.Get("/api/v1/users", handlers.GetUsersHandler)
	})
	r.Post("/api/v1/users", handlers.CreateUserHandler)
	r.Delete("/api/v1/users", handlers.DeleteUserHandler)

	r.Get("/api/v1/locations", handlers.GetLocationsHandler)
	r.Post("/api/v1/locations", handlers.CreateLocationHandler)
	r.Delete("/api/v1/locations", handlers.DeleteLocationHandler)

	slog.Info("Starting chi server on port 8080!")
	http.ListenAndServe(":8080", r)

}
