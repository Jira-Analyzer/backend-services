package jira

import "time"

type Config struct {
	JiraUrl          string        `yaml:"jiraUrl"`
	IssuesPerRequest int           `yaml:"issuesPerRequest"`
	ThreadCount      int           `yaml:"threadCount"`
	MaxTimeSleep     time.Duration `yaml:"maxTimeSleep"`
	MinTimeSleep     time.Duration `yaml:"minTimeSleep"`
}
