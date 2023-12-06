package jira

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type ConnectorServiceConfig struct {
	LocalPort        int64         `yaml:"localHttpServerPort"`
	JiraUrl          string        `yaml:"jiraUrl"`
	IssuesPerRequest int64         `yaml:"issuesPerRequest"`
	ThreadCount      int64         `yaml:"threadCount"`
	MaxTimeSleep     time.Duration `yaml:"maxTimeSleep"`
	MinTimeSleep     time.Duration `yaml:"minTimeSleep"`
}

type Config struct {
	ConnectorConfig ConnectorServiceConfig `yaml:"connector"`
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
