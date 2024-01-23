package config

import (
	"fmt"
	"os"
	"time"

	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	"github.com/Jira-Analyzer/backend-services/internal/logger"
	"github.com/Jira-Analyzer/backend-services/internal/server"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Server defaults
const (
	defaultReadTimeout      time.Duration = 5 * time.Second
	defaultWriteTimeout     time.Duration = 5 * time.Second
	defaultServerHost       string        = ":8000"
	defaultResourseTimeout  time.Duration = 5 * time.Second
	defaultAnalyticsTimeout time.Duration = 15 * time.Second
)

type Config struct {
	DbConfig     provider.DbConfig   `yaml:"db"`
	ServerConfig server.ServerConfig `yaml:"rest-service"`
	LoggerConfig logger.LoggerConfig `yaml:"logger"`
}

func (config *Config) ValidateConfig() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(config)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		logrus.Info(errors)
		return err
	}
	return nil
}

func (config *Config) PopulateConfig() {
	s := func(s string) *string { return &s }
	d := func(d time.Duration) *time.Duration { return &d }

	if config.ServerConfig.Host == nil {
		config.ServerConfig.Host = s(defaultServerHost)
	}
	if config.ServerConfig.AnalyticsTimeout == nil {
		config.ServerConfig.AnalyticsTimeout = d(defaultAnalyticsTimeout)
	}
	if config.ServerConfig.ResourceTimeout == nil {
		config.ServerConfig.ReadTimeout = d(defaultResourseTimeout)
	}
	if config.ServerConfig.ReadTimeout == nil {
		config.ServerConfig.ReadTimeout = d(defaultReadTimeout)
	}
	if config.ServerConfig.WriteTimeout == nil {
		config.ServerConfig.WriteTimeout = d(defaultWriteTimeout)
	}
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
