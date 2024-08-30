package handlers

import (
	"file-registory/responses"
	s "file-registory/server"

	"github.com/gofiber/fiber/v2"
)

type HandlerFiles struct {
	Server *s.Server
}

func NewHandlerFiles(server *s.Server) *HandlerFiles {
	return &HandlerFiles{
		Server: server,
	}
}

// @Summary Save file
// @Description Uploads it to the IPFS node using one of the client libraries. IPFS client will return the cid for the uploaded file. Saves the data in the File Registry.
// @Tags Files
// @Accept json
// @Produce json
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /files [post]
func (h *HandlerFiles) SaveFile(c *fiber.Ctx) error {
	return responses.MessageResponse(c, fiber.StatusOK, "Hellow World")
}

// @Summary Get file
// @Description Get the data from the File Registry and return the CID to the client.
// @Tags Files
// @Accept json
// @Produce json
// @Param filePath query string true "File Path"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /files [get]
func (h *HandlerFiles) GetFile(c *fiber.Ctx) error {
	return responses.MessageResponse(c, fiber.StatusOK, "Hellow World")
}
