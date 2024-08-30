package main

import (
	cfg "file-registory/pkgs/config"
	s "file-registory/server"
	"file-registory/server/routes"
	"log"
)

// @Title File Registry
// @Version 1.0
// @description This is a API for File Registry.
// @BasePath /api/v1
func main() {
	config := cfg.NewConfig()

	err := config.LoadEnvironment()
	if err != nil {
		log.Fatal("Failed to load environment variables!")
	}

	server := s.NewServer(config)

	routes.ConfigureRoutes(server)
	server.Listen()
}
