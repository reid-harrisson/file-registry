package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string
	IpfsUrl      string
	RpcSrvUrl    string
	PrivateKey   string
	ContractAddr string
}

func NewConfig() *Config {
	return &Config{
		ServerPort: "8000", // Default Server Port
	}
}

func (config *Config) LoadEnvironment() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	config.ServerPort = os.Getenv("SERVER_PORT")
	config.IpfsUrl = os.Getenv("IPFS_URL")
	config.RpcSrvUrl = os.Getenv("RPC_SRV_URL")
	config.PrivateKey = os.Getenv("PRIVATE_KEY")
	config.ContractAddr = os.Getenv("CONTRACT_ADDR")

	return nil
}
