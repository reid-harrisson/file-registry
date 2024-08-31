package main

import (
	cfg "file-registory/config"
	s "file-registory/server"
	"file-registory/server/routes"
	"log"
)

// @Title File Registry
// @Version 1.0
// @description This is a API for File Registry.
func main() {
	config := cfg.NewConfig()

	err := config.LoadEnvironment()
	if err != nil {
		log.Println("Failed to load environment variables!")
		log.Println("Server port is set to default(8000).")
	}

	server := s.NewServer(config)

	routes.ConfigureRoutes(server)
	server.Listen()
}
