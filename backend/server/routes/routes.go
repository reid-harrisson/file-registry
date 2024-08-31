package routes

import (
	s "file-registory/server"
	"file-registory/server/handlers"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/swaggo/fiber-swagger"
)

func ConfigureRoutes(server *s.Server) {
	server.App.Get("/docs/*", swagger.WrapHandler)

	v1 := server.App.Group("/v1")

	groupFiles := v1.Group("/files")
	GroupFiles(server, groupFiles)

}

func GroupFiles(server *s.Server, router fiber.Router) {
	handler := handlers.NewHandlerFiles(server)

	router.Get("", handler.GetFile)
	router.Post("", handler.GetFile)
}
