package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
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
	if config.ServerPort == "" {
		config.ServerPort = "8000"
	}

	return nil
}
