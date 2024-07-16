package configs

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App      App
		Database Database
		Mongo    Mongo
	}
)

// New returns app config.
func New() (*Config, error) {
	config := &Config{}

	if _, err := os.Stat("./env/.env"); err == nil {
		err = cleanenv.ReadConfig("./env/.env", config)
		if err != nil {
			return config, fmt.Errorf("config error: %w", err)
		}
	}

	err := cleanenv.ReadEnv(config)
	if err != nil {
		return config, err
	}

	return config, nil
}
