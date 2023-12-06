package connector

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

type ConnectorServiceConfig struct {
	Local_Port         int64         `yaml:"local_http_server_port"`
	JiraUrl            string        `yaml:"jiraUrl"`
	Issues_Per_Request int64         `yaml:"issuesPerRequest"`
	Thread_Count       int64         `yaml:"threadCount"`
	Max_Time_Sleep     time.Duration `yaml:"maxTimeSleep"`
	MinTimeSleep       time.Duration `yaml:"minTimeSleep"`
}

type Config struct {
	DbConfig        DbConfig               `yaml:"db"`
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
