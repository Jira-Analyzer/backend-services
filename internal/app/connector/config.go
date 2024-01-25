package connector

import (
	"github.com/Jira-Analyzer/backend-services/internal/client/jira"
	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	"github.com/Jira-Analyzer/backend-services/internal/logger"
	"github.com/Jira-Analyzer/backend-services/internal/server"
)

type Config struct {
	DbConfig        *provider.DbConfig   `yaml:"dbConfig"`
	LoggerConfig    *logger.LoggerConfig `yaml:"loggerConfig"`
	ServerConfig    *server.ServerConfig `yaml:"serverConfig"`
	ConnectorConfig *jira.Config         `yaml:"connector"`
}
