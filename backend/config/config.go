package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort      string
	IpfsUrl         string
	SepoliaUrl      string
	PrivateKey      string
	ContractAddress string
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
	config.SepoliaUrl = os.Getenv("SEPOLIA_URL")
	config.PrivateKey = os.Getenv("PRIVATE_KEY")
	config.ContractAddress = os.Getenv("CONTRACT_ADDRESS")

	return nil
}
