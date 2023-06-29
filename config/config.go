package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type (
	Config struct {
		Logger           Logger           `yaml:"logger"`
		TgAPI            TelegramAPI      `yaml:"telegram_api"`
		MongoDB          Repository       `yaml:"mongo-db"`
		CurrencyRatesAPI CurrencyRatesAPI `yaml:"currency_rates"`
	}

	Logger struct {
		Level string `yaml:"level"`
	}

	TelegramAPI struct {
		AccessToken string `yaml:"access_token"`
	}

	CurrencyRatesAPI struct {
		AccessToken string `yaml:"access_token"`
	}

	Repository struct {
		URI      string `yaml:"uri"`
		DataBase string `yaml:"database"`
	}
)

func NewConfig(fileName string) (*Config, error) {
	f, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var cfg *Config
	if err = yaml.Unmarshal(f, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
