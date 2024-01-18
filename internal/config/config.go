package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type RESTServiceConfig struct {
	Host             string        `yaml:"host"`
	ResourceTimeout  time.Duration `yaml:"resource-timeout"`
	AnalyticsTimeout time.Duration `yaml:"analytics-timeout"`
}

type Config struct {
	DbConfig          DbConfig          `yaml:"db"`
	RESTServiceConfig RESTServiceConfig `yaml:"rest-service"`
}

func ReadConfigFromYAML(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file from %s, due to: %w", path, err)
	}
	conf := &Config{}
	err = yaml.Unmarshal(file, conf)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config from %s, due to: %w", path, err)
	}

	return conf, nil
}
