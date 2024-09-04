package server

import (
	"file-registory/config"
	"file-registory/contracts"
	_ "file-registory/docs"

	"github.com/gofiber/fiber/v2"
	shell "github.com/ipfs/go-ipfs-api"
)

type Server struct {
	App      *fiber.App
	Config   *config.Config
	Contract *contracts.Contract
	IPFS     *shell.Shell
}

func NewServer(config *config.Config) (*Server, error) {
	app := fiber.New()

	contract, err := contracts.NewContract(config)
	if err != nil {
		return nil, err
	}

	sh := shell.NewShell(config.IpfsUrl)

	return &Server{
		App:      app,
		Config:   config,
		Contract: contract,
		IPFS:     sh,
	}, nil
}

func (server *Server) Listen() {
	server.App.Listen(":8000")
}
