package jira

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetchIssues(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
			"total": 1,
			"issues": [{
				"key": "ISSUE-123",
				"fields": {
					"summary": "Issue Summary",
					"issuetype": {
						"name": "Task"
					},
					"status": {
						"name": "Open"
					},
					"priority": {
						"name": "High"
					},
					"creator": {
						"name": "John Doe"
					},
					"project": {
						"name": "Sample Project"
					},
					"description": "Issue Description",
					"assignee": {
						"name": "Jane Doe"
					},
					"created": "2022-01-01T12:00:00Z",
					"updated": "2022-01-02T14:30:00Z",
					"resolutiondate": "2022-01-03T10:15:00Z"
				}
			}]
		}`))
	}))
	defer server.Close()

	client := &Client{
		baseURL:          server.URL,
		client:           &http.Client{},
		issuesPerRequest: 1,
		threadCount:      1,
		maxTimeSleep:     time.Second,
		minTimeSleep:     time.Millisecond,
	}

	issues, err := client.FetchIssues("SampleProject", 100)

	assert.NoError(t, err)
	assert.Len(t, issues, 1)
	assert.Contains(t, issues, Issue{
		Key: "ISSUE-123",
		Fields: IssueFields{
			Summary: "Issue Summary",
			Type: struct {
				Name string `json:"name"`
			}{Name: "Task"},
			Status: struct {
				Name string `json:"name"`
			}{Name: "Open"},
			Priority: struct {
				Name string `json:"name"`
			}{Name: "High"},
			Creator: struct {
				Name string `json:"name"`
			}{Name: "John Doe"},
			Project: struct {
				Name string `json:"name"`
			}{Name: "Sample Project"},
			Description: "Issue Description",
			AssigneeName: struct {
				Name string `json:"name"`
			}{Name: "Jane Doe"},
			CreatedTime: time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC),
			UpdatedTime: time.Date(2022, 1, 2, 14, 30, 0, 0, time.UTC),
			ClosedTime:  time.Date(2022, 1, 3, 10, 15, 0, 0, time.UTC),
		},
	})
}

func TestFetchProjects(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`[
			{
				"name": "Project1",
				"self": "http://example.com/projects/project1"
			},
			{
				"name": "Project2",
				"self": "http://example.com/projects/project2"
			}
		]`))
	}))
	defer server.Close()

	client := &Client{
		baseURL:          server.URL,
		client:           &http.Client{},
		issuesPerRequest: 1,
		threadCount:      1,
		maxTimeSleep:     time.Second,
		minTimeSleep:     time.Millisecond,
	}

	projectsResponse, err := client.FetchProjects(10, 1, "Project", "project1")

	assert.NoError(t, err)
	assert.Len(t, projectsResponse.Projects, 2)
	assert.Equal(t, "Project1", projectsResponse.Projects[0].Name)
	assert.Equal(t, "http://example.com/projects/project1", projectsResponse.Projects[0].Link)
	assert.Equal(t, "Project2", projectsResponse.Projects[1].Name)
	assert.Equal(t, "http://example.com/projects/project2", projectsResponse.Projects[1].Link)
	assert.Equal(t, 1, projectsResponse.PageInfo.PageCount)
	assert.Equal(t, 1, projectsResponse.PageInfo.CurrentPage)
	assert.Equal(t, 2, projectsResponse.PageInfo.ProjectsCount)
}
