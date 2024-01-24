package connector

import (
	"fmt"
	"github.com/Jira-Analyzer/backend-services/internal/client/jira"
	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	"github.com/Jira-Analyzer/backend-services/internal/logger"
	"github.com/Jira-Analyzer/backend-services/internal/server"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DbConfig        *provider.DbConfig   `yaml:"dbConfig"`
	LoggerConfig    *logger.LoggerConfig `yaml:"loggerConfig"`
	ServerConfig    *server.ServerConfig `yaml:"serverConfig"`
	ConnectorConfig *jira.Config         `yaml:"connector"`
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
