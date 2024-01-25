package jira

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createMockServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/rest/api/2/project/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"id": "1", "name": "Sample Project", "description": "Sample Description"}`))
	})

	handler.HandleFunc("/rest/api/2/project", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[{"name": "Project1", "link": "link1"}, {"name": "Project2", "link": "link2"}]`))
	})

	handler.HandleFunc("/rest/api/2/search", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"total": 1, "issues": [{"id": "2", "key": "ABC-123", "fields": {"summary": "Issue Summary"}}]}`))
	})

	return httptest.NewServer(handler)
}

func TestFetchProject(t *testing.T) {
	mockServer := createMockServer()
	defer mockServer.Close()

	config := &Config{JiraUrl: mockServer.URL}
	client := NewClient(config)

	project, err := client.FetchProject(1)

	assert.NoError(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, "Sample Project", project.Name)
	assert.Equal(t, "Sample Description", project.Description)
}

func TestFetchProjects(t *testing.T) {
	mockServer := createMockServer()
	defer mockServer.Close()

	config := &Config{JiraUrl: mockServer.URL}
	client := NewClient(config)

	projectsResponse, err := client.FetchProjects(1, 10)

	assert.NoError(t, err)
	assert.NotNil(t, projectsResponse)
	assert.Equal(t, 2, len(projectsResponse.Projects))
	assert.Equal(t, 1, projectsResponse.PageInfo.PageCount)
}

func TestFetchIssues(t *testing.T) {
	mockServer := createMockServer()
	defer mockServer.Close()

	config := &Config{JiraUrl: mockServer.URL, ThreadCount: 2, IssuesPerRequest: 1, MaxTimeSleep: 500, MinTimeSleep: 100}
	client := NewClient(config)

	issues, err := client.FetchIssues(1, 100)

	assert.NoError(t, err)
	assert.NotEmpty(t, issues)

	keyFound := false
	for _, issue := range issues {
		if issue.Key == "ABC-123" {
			keyFound = true
			break
		}
	}

	assert.True(t, keyFound, "Ожидался ключ ABC-123 в полученных данных")
}
