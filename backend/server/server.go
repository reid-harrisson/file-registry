package server

import (
	"file-registory/config"
	_ "file-registory/docs"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App    *fiber.App
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	app := fiber.New()
	return &Server{
		App:    app,
		Config: config,
	}
}

func (server *Server) Listen() {
	server.App.Listen(":8000")
}
