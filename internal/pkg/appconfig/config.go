package appconfig

import (
	"errors"
	"os"
)

var (
	ErrMissingAPIToken = errors.New("missing telegram api token")
)

type Config struct {
	TelegramAPIToken string
}

func loadConfig() Config {
	return Config{
		TelegramAPIToken: os.Getenv("TELEGRAM_API_TOKEN"),
	}
}

func (c Config) validate() error {
	if c.TelegramAPIToken == "" {
		return ErrMissingAPIToken
	}

	return nil
}
