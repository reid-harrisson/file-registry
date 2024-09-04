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
// @Tags files
// @Accept multipart/form-data
// @Produce json
// @Param filePath query string true "File Path"
// @Param file formData file true "File to Upload"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/files [post]
func (h *HandlerFiles) SaveFile(c *fiber.Ctx) error {
	filePath := c.Query("filePath")

	file, err := c.FormFile("file")
	if err != nil {
		responses.ErrorResponse(c, fiber.StatusBadRequest, "Failed to upload file.")
	}

	src, err := file.Open()
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, "Failed to upload file.")
	}
	defer src.Close()

	cid, err := h.Server.IPFS.Add(src)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, "Failed to upload file.")
	}

	err = h.Server.Contract.Save(filePath, cid)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.MessageResponse(c, fiber.StatusOK, "Sucessfully saved.")
}

// @Summary Get file
// @Description Get the data from the File Registry and return the CID to the client.
// @Tags files
// @Accept json
// @Produce json
// @Param filePath query string true "File Path"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /v1/files [get]
func (h *HandlerFiles) GetFile(c *fiber.Ctx) error {
	filePath := c.Query("filePath")

	cid, err := h.Server.Contract.Get(filePath)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.MessageResponse(c, fiber.StatusOK, cid)
}
