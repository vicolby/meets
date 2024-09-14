package tests

import (
	"log"
	"os"
	"testing"

	"github.com/vicolby/events/database"
	"github.com/vicolby/events/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	connString, terminate, err := setupTestContainer()
	if err != nil {
		log.Fatalf("could not set up test container: %v", err)
	}
	defer terminate()

	log.Print("Started test postgres container")

	database.DB, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if database.DB == nil {
		log.Fatalf("database connection is nil")
	}

	database.DB.AutoMigrate(types.User{}, types.Event{}, types.Location{})

	code := m.Run()
	os.Exit(code)
}
