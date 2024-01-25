package jira

import "time"

type Config struct {
	Host             string        `yaml:"host"`
	JiraUrl          string        `yaml:"jira-url"`
	IssuesPerRequest int           `yaml:"issues-per-request"`
	ThreadCount      int           `yaml:"thread-count"`
	MaxTimeSleep     time.Duration `yaml:"max-time-sleep"`
	MinTimeSleep     time.Duration `yaml:"min-time-sleep"`
	ReadTimeout      time.Duration `yaml:"read-timeout"`
	WriteTimeout     time.Duration `yaml:"write-timeout"`
}
