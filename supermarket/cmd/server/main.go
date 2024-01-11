package main

import (
	"fmt"
	"os"
	"supermarket/internal/application"
)

func main() {
	os.Setenv("Token", "themostsecrettoken")
	// create and start server
	server := application.NewServer(":8080")
	if err := server.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
