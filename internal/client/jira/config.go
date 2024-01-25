package jira

type Config struct {
	JiraUrl          string `yaml:"jiraUrl"`
	IssuesPerRequest int    `yaml:"issuesPerRequest"`
	ThreadCount      int    `yaml:"threadCount"`
	MaxTimeSleep     int    `yaml:"maxTimeSleep"`
	MinTimeSleep     int    `yaml:"minTimeSleep"`
}
