package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	SecretId  string   `yaml:"secret_id"`
	SecretKey string   `yaml:"secret_key"`
	CertPath  string   `yaml:"cert_path"`
	KeyPath   string   `yaml:"key_path"`
	Domains   []string `yaml:"domains"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
