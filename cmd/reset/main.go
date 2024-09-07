package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vicolby/events/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	db, err := database.ConnectToPostgres()
	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()

	tables := []string{
		"schema_migrations",
		"locations",
		"events",
		"users",
	}

	for _, table := range tables {
		query := fmt.Sprintf("drop table if exists %s cascade", table)
		if _, err := sqlDB.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
}
