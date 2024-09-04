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
		log.Fatal(err)
	}

	server, err := s.NewServer(config)

	if err != nil {
		log.Fatal(err)
	}

	routes.ConfigureRoutes(server)
	server.Listen()
}
