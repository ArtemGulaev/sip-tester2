package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	SIP SIPConfig `yaml:"sip"`
}

type SIPConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Domain   string `yaml:"domain"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(data, cfg)

	return cfg, err
}
