package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	GitHubToken string `env:"GITHUB_TOKEN"`
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	config := &Config{}
	err := env.Parse(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
