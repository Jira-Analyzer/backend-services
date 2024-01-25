package main

import (
	"fmt"
	"github.com/Jira-Analyzer/backend-services/internal/client/jira"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func main() {
	//https://issues.apache.org/jira/rest/api/2/project
	config := jira.Config{
		JiraUrl:          "https://issues.apache.org/jira",
		IssuesPerRequest: 10,
		ThreadCount:      5,
		MaxTimeSleep:     time.Second * 5,
		MinTimeSleep:     time.Second * 2,
	}

	// Создание клиента Jira
	client := jira.NewClient(config)

	// Пример вызова FetchProjects
	projects, err := client.FetchProjects()

	if err != nil {
		logrus.Error(err)
	}
	fmt.Printf("Fetched projects: %+v\n", projects)

	// Пример вызова FetchIssues (предполагается, что у вас есть ключ проекта)
	projectKey := "12345"
	issues, err := client.FetchIssues(projectKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Fetched issues for project %s: %+v\n", projectKey, issues)
}
