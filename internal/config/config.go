package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	HTTP Server   `yaml:"server"`
	DB   Database `yaml:"database"`
}

type Server struct {
	Port      string `yaml:"port"`
	JWTSecret string `yaml:"jwt_secret"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Basename string `yaml:"basename"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func Load(configPath string) (*Config, error) {
	cfg := &Config{}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(file, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return cfg, nil
}
