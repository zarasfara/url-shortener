package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP httpConfig `yaml:"http"`
		DB   dbConfig   `yaml:"db"`
	}

	httpConfig struct {
		Address string `yaml:"address"`
		Port    string `yaml:"port"`
	}

	dbConfig struct {
		Database string `env:"DB_DATABASE"`
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		Username string `env:"DB_USERNAME"`
		Password string `env:"DB_PASSWORD"`
		SSLMode  string `yaml:"sslmode"`
	}
)

func MustLoad(env string) *Config {
	var cfg Config

	configPath := fmt.Sprintf("./configs/%s.yml", env)

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("can't read config: %s", err)
	}

	return &cfg
}
