package main

import (
	"fmt"
	"os"
	"supermarket/internal/application"
)

func main() {
	// server config
	config := application.ServerConfig{
		Host:   os.Getenv("ENV_HOST"),
		Port:   os.Getenv("ENV_PORT"),
		DbFile: os.Getenv("ENV_PATH_DBFILE"),
		Token:  os.Getenv("ENV_TOKEN"),
	}
	// create and start server
	server := application.NewServer(config)
	if err := server.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
