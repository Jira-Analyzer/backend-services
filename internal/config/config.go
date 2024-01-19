package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	Host     *string `yaml:"host"`
	Name     *string `yaml:"name"`
	User     *string `yaml:"user"`
	Password *string `yaml:"password"`
}

type ServerConfig struct {
	Host             *string        `yaml:"host"`
	ResourceTimeout  *time.Duration `yaml:"resource-timeout"`
	AnalyticsTimeout *time.Duration `yaml:"analytics-timeout"`
}

type Config struct {
	DbConfig     DbConfig     `yaml:"db"`
	ServerConfig ServerConfig `yaml:"rest-service"`
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

func ValidateConfig(config *Config) error {
	var err error = nil
	fields := make([]string, 0)
	if config.DbConfig.Host == nil {
		fields = append(fields, "db.Host")
	}
	if config.DbConfig.Name == nil {
		fields = append(fields, "db.Name")
	}
	if config.DbConfig.User == nil {
		fields = append(fields, "db.User")
	}
	if config.DbConfig.Password == nil {
		fields = append(fields, "db.Password")
	}

	if len(fields) != 0 {
		err = fmt.Errorf("Error: %w. Missing fields %s", errorlib.ConfigValueMissError, strings.Join(fields, ", "))
	}
	return err
}
