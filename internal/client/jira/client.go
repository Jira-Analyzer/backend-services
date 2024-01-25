package jira

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Client struct {
	baseURL          string
	client           *http.Client
	issuesPerRequest int
	threadCount      int
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

func (c *Client) FetchProjects(limit int, page int, search string, projectKey string) (*ProjectsResponse, error) {

	response, err := http.Get(c.baseURL + "/rest/api/2/project")
	if err != nil {
		logrus.Error("Unable to get projects list ")
		return &ProjectsResponse{}, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return &ProjectsResponse{}, err
	}

	var jiraProjects []JiraProject
	err = json.Unmarshal(body, &jiraProjects)

	if err != nil {
		return &ProjectsResponse{}, err
	}

	var projects []Project

	projectsCount := 0

	for _, elem := range jiraProjects {
		if strings.Contains(strings.ToLower(projectKey), strings.ToLower(search)) {
			projectsCount++
			projects = append(projects, Project{
				Name: elem.Name,
				Link: elem.Link,
			})
		}
	}

	startIndex := limit * (page - 1)
	endIndex := startIndex + limit
	if endIndex >= len(projects) {
		endIndex = len(projects)
	}

	return &ProjectsResponse{
		Projects: projects[startIndex:endIndex],
		PageInfo: PageInfo{
			PageCount:     int(math.Ceil(float64(projectsCount) / float64(limit))),
			CurrentPage:   page,
			ProjectsCount: projectsCount,
		},
	}, nil
}

func (c *Client) FetchIssues(projectKey string, timeToWaitMs int) (map[Issue]struct{}, error) {
	httpClient := &http.Client{}
	response, err := httpClient.Get(c.baseURL +
		"/rest/api/2/search?jql=project=" + projectKey + "&expand=changelog&startAt=0&maxResults=1")

	if err != nil || response.StatusCode != http.StatusOK {
		logrus.Error("Unable to get issues for project " + projectKey)
		return map[Issue]struct{}{}, nil
	}

	body, _ := io.ReadAll(response.Body)
	var issueResponse IssuesList
	if err = json.Unmarshal(body, &issueResponse); err != nil {
		logrus.Error("Error while unmarshalling issue response")
	}

	totalIssuesCount := issueResponse.IssuesCount

	if totalIssuesCount == 0 {
		return map[Issue]struct{}{}, nil
	}

	issues := map[Issue]struct{}{}
	issues[issueResponse.Issues[0]] = struct{}{}

	waitGroup := sync.WaitGroup{}
	mutex := sync.Mutex{}
	wasError := false

	threadCount := c.threadCount
	issuesPerRequest := c.issuesPerRequest

	stop := make(chan struct{})

	for i := 0; i < threadCount; i++ {
		waitGroup.Add(1)
		go func(threadNumber int) {
			defer waitGroup.Done()
			select {
			case <-stop:
				logrus.Error("Error while reading issues in thread... Stopping all other threads...")
				return
			default:
				threadStartIndex := (totalIssuesCount/threadCount)*threadNumber + 1
				requestCount := int(math.Ceil(float64(totalIssuesCount) / float64(threadCount*issuesPerRequest)))
				for j := 0; j < requestCount; j++ {
					startAt := threadStartIndex + j*issuesPerRequest
					if startAt < totalIssuesCount {
						requestString := c.baseURL + "/rest/api/2/search?jql=project=" +
							projectKey + "&expand=changelog&startAt=" + strconv.Itoa(startAt) +
							"&maxResults=" + strconv.Itoa(issuesPerRequest)

						response, requestErr := httpClient.Get(requestString)
						body, responseReadErr := io.ReadAll(response.Body)

						if requestErr != nil || responseReadErr != nil {
							wasError = true
							close(stop)
							return
						}

						var issueResponse IssuesList
						_ = json.Unmarshal(body, &issueResponse)

						mutex.Lock()
						for _, elem := range issueResponse.Issues {
							issues[elem] = struct{}{}
						}
						mutex.Unlock()
					}
				}
			}
		}(i)
	}
	waitGroup.Wait()

	if wasError {
		time.Sleep(time.Duration(timeToWaitMs) * time.Millisecond)
		newTimeToSleep := int(math.Ceil(float64(timeToWaitMs) * math.Phi))
		logrus.Error("Error while downloading issues for project \"" +
			projectKey + "\", waiting now" + strconv.Itoa(timeToWaitMs) + "ms")

		if time.Duration(newTimeToSleep) > c.maxTimeSleep {
			return map[Issue]struct{}{}, errors.New("A lot of time to sleep")
		}

		return c.FetchIssues(projectKey, newTimeToSleep)
	}

	return issues, nil
}
