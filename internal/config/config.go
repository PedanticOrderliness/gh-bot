package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Schema          string `yaml:"schema" env-required:"true"`
	BotAccessToken  string `env:"BOT_ACCESS_TOKEN" env-required:"true"`
	Port            int    `yaml:"port" env:"PORT" env-default:"8080"`
	GhAppID         int64  `yaml:"gh_app_id" env:"GH_APP_ID" env-default:"245647"`
	GhInstllationID int64  `yaml:"gh_installation_id" env:"GH_INSTALLATION_ID" env-default:"30283851"`
}

func GetConfig(filename string) (*Config, error) {
	config := &Config{}
	err := cleanenv.ReadConfig(filename, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
