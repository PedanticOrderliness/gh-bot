package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Schema         string `yaml:"schema" env-required:"true"`
	BotAccessToken string `env:"BOT_ACCESS_TOKEN" env-required:"true"`
	Port           int    `yaml:"port" env:"PORT" env-default:"80"`
}

func GetConfig(filename string) (*Config, error) {
	config := &Config{}
	err := cleanenv.ReadConfig(filename, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
