package jira

import (
	"time"
)

type Config struct {
	JiraUrl          string        `yaml:"jiraUrl"`
	IssuesPerRequest int64         `yaml:"issuesPerRequest"`
	ThreadCount      int64         `yaml:"threadCount"`
	MaxTimeSleep     time.Duration `yaml:"maxTimeSleep"`
	MinTimeSleep     time.Duration `yaml:"minTimeSleep"`
}
