package main

import (
	"fmt"

	"github.com/vicolby/meets/server"
)

func main() {
	server := server.NewServer(":3000")
	fmt.Println("Server started on port: 3000")
	server.Start()
}
