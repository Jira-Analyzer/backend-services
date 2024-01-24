package jira

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseURL          string
	client           *http.Client
	issuesPerRequest int64
	threadCount      int64
	maxTimeSleep     time.Duration
	minTimeSleep     time.Duration
}

func NewClient(config Config) *Client {
	return &Client{
		baseURL:          config.JiraUrl,
		client:           &http.Client{},
		issuesPerRequest: config.IssuesPerRequest,
		threadCount:      config.ThreadCount,
		maxTimeSleep:     config.MaxTimeSleep,
		minTimeSleep:     config.MinTimeSleep,
	}
}

func (c *Client) FetchProjects() (*FetchProjectResponseDTO, error) {
	url := fmt.Sprintf("%s/rest/api/2/project", c.baseURL)

	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var dto FetchProjectResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (c *Client) FetchIssues(projectKey string) (*FetchIssueResponseDTO, error) {
	url := fmt.Sprintf("%s/rest/api/2/search?jql=project=%s", c.baseURL, projectKey)

	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var dto FetchIssueResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}
