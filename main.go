package main

import (
	"fmt"

	"github.com/vicolby/meets/server"
	"github.com/vicolby/meets/storer"
)

func main() {
	conn, err := storer.CreateConnection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v \n", err)
		panic(err)
	}
	defer conn.Close()

	storage := storer.NewPostgresStorage(conn)
	err = storage.Init()

	if err != nil {
		fmt.Printf("Error initializing database: %v \n", err)
		panic(err)
	}

	userStorage := storer.NewUserPostgresStorage(conn)
	eventStorage := storer.NewEventPostgresStorage(conn)

	server := server.NewServer(
		"localhost:3000",
		userStorage,
		eventStorage,
	)
	fmt.Println("Server started on port: 3000")
	server.Start()
}
