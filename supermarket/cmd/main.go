package main

import "supermarket/internal/application"

func main() {
	// create and start server
	server := application.NewServer(":8080")
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
