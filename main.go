package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/vicolby/events/database"
	"github.com/vicolby/events/handlers"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger
var validate *validator.Validate

func main() {
	l, err := zap.NewProduction()
	logger = l.Sugar()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	err = godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}

	err = database.Init()

	if err != nil {
		logger.Info("Error connecting to database:", zap.Error(err))
		return
	}

	logger.Info("Successfully connected to the database!")

	validate = validator.New(validator.WithRequiredStructEnabled())

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)

	eventHandler := handlers.NewEventHandler(logger, validate)
	userHandler := handlers.NewUserHandler(logger, validate)
	locationHandler := handlers.NewLocationHandler(logger, validate)

	r.Get("/api/v1/events", eventHandler.GetEvents)
	r.Post("/api/v1/events", eventHandler.CreateEvent)
	r.Delete("/api/v1/events", eventHandler.DeleteEvent)

	r.Get("/api/v1/users", userHandler.GetUsers)
	r.Post("/api/v1/users", userHandler.CreateUser)
	r.Delete("/api/v1/users", userHandler.DeleteUser)

	r.Get("/api/v1/locations", locationHandler.GetLocations)
	r.Post("/api/v1/locations", locationHandler.CreateLocation)
	r.Delete("/api/v1/locations", locationHandler.DeleteLocation)

	// r.Group(func(r chi.Router) {
	// 	r.Use(handlers.EnsureValidToken())

	// 	r.Get("/api/v1/events", handlers.GetEventsHandler)
	// 	r.Post("/api/v1/events", handlers.CreateEventHandler)
	// 	r.Delete("/api/v1/events", handlers.DeleteEventHandler)
	// 	r.Put("/api/v1/events", handlers.UpdateEventHandler)
	// 	r.Put("/api/v1/events/{eventID}/add_participants", handlers.AddEventParticipant)
	// 	r.Delete("/api/v1/events/{eventID}/delete_participant", handlers.DeleteEventParticipant)

	// 	r.Get("/api/v1/users", handlers.GetUsersHandler)
	// 	r.Post("/api/v1/users", handlers.CreateUserHandler)
	// 	r.Delete("/api/v1/users", handlers.DeleteUserHandler)

	// 	r.Get("/api/v1/locations", handlers.GetLocationsHandler)
	// 	r.Post("/api/v1/locations", handlers.CreateLocationHandler)
	// 	r.Delete("/api/v1/locations", handlers.DeleteLocationHandler)

	// })

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("HTTP server ListenAndServe:", zap.Error(err))
		}
	}()
	logger.Info("Server is running on port 8080")

	gracefulShutdown(srv, 5*time.Second)
}

func gracefulShutdown(srv *http.Server, timeout time.Duration) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	logger.Infof("Received signal: %v. Initiating graceful shutdown...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	cleanupResources()

	logger.Info("Server shutdown gracefully")
}

func cleanupResources() {
	if database.DB != nil {
		sqlDB, err := database.DB.DB()
		if err != nil {
			logger.Infof("Error getting SQL DB instance: %v", err)
			return
		}

		if err := sqlDB.Close(); err != nil {
			logger.Infof("Error closing the database: %v", err)
		} else {
			logger.Info("Database connection closed")
		}
	}

	logger.Info("Logs flushed")
}
