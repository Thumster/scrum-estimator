package appconfig

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

var (
	ErrLoadingEnvFile = errors.New("error loading .env file")
)

func LoadFromEnv() (Config, error) {
	if os.Getenv("APP_ENV") != "prod" {
		if err := godotenv.Load(); err != nil {
			return Config{}, ErrLoadingEnvFile
		}
	}
	c := loadConfig()

	if err := c.validate(); err != nil {
		return c, err
	}

	return c, nil
}
