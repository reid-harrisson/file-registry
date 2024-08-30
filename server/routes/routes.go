package routes

import (
	s "file-registory/server"
	"file-registory/server/handlers"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/swaggo/fiber-swagger"
)

func ConfigureRoutes(server *s.Server) {
	server.App.Get("/docs/*", swagger.WrapHandler)

	apiV1 := server.App.Group("/api/v1")

	groupFiles := apiV1.Group("/files")
	GroupFiles(server, groupFiles)

}

func GroupFiles(server *s.Server, router fiber.Router) {
	handler := handlers.NewHandlerFiles(server)

	router.Get("", handler.GetFile)
	router.Post("", handler.GetFile)
}
