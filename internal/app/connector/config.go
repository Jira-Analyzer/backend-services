package connector

import (
	"github.com/Jira-Analyzer/backend-services/internal/client/jira"
	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	"github.com/Jira-Analyzer/backend-services/internal/logger"
)

type Config struct {
	DbConfig        *provider.DbConfig   `yaml:"db"`
	LoggerConfig    *logger.LoggerConfig `yaml:"logger"`
	ConnectorConfig *jira.Config         `yaml:"connector-service"`
}
