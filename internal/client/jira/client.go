package jira

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/Jira-Analyzer/backend-services/internal/client/jira/dto"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
	"github.com/sirupsen/logrus"
)

type Client struct {
	baseURL          string
	client           *http.Client
	issuesPerRequest int
	threadCount      int
	maxTimeSleep     time.Duration
	minTimeSleep     time.Duration
}

func NewClient(config *Config) *Client {
	return &Client{
		baseURL:          config.JiraUrl,
		client:           &http.Client{},
		issuesPerRequest: config.IssuesPerRequest,
		threadCount:      config.ThreadCount,
		maxTimeSleep:     config.MaxTimeSleep,
		minTimeSleep:     config.MinTimeSleep,
	}
}

func (c *Client) FetchProject(id int) (*domain.Project, error) {
	key := strconv.Itoa(id)
	response, err := http.Get(c.baseURL + "/rest/api/2/project/" + key)
	if err != nil {
		logrus.Error("Unable to get project details")
		return nil, err
	}

	var jiraProject dto.JiraProject
	if err := json.NewDecoder(response.Body).Decode(&jiraProject); err != nil {
		return nil, err
	}

	project := &domain.Project{
		Name:        jiraProject.Name,
		Id:          id,
		Description: jiraProject.Description,
		AvatarUrl:   jiraProject.AvatarUrls.Url,
		Type:        jiraProject.Type,
		Archived:    jiraProject.Archived,
	}

	return project, nil
}

func (c *Client) FetchProjects(page int, count int) (*dto.ProjectsResponse, error) {
	response, err := http.Get(c.baseURL + "/rest/api/2/project")
	if err != nil {
		logrus.Error("Unable to get projects list")
		return &dto.ProjectsResponse{}, err
	}

	var jiraProjects []dto.JiraProject
	if err := json.NewDecoder(response.Body).Decode(&jiraProjects); err != nil {
		return &dto.ProjectsResponse{}, err
	}

	projects := make([]dto.Project, 0)
	for _, elem := range jiraProjects {
		projects = append(projects, dto.Project{
			Name: elem.Name,
			Link: elem.Link,
		})
	}

	startIndex := count * (page - 1)
	endIndex := startIndex + count
	if endIndex > len(projects) {
		endIndex = len(projects)
	}
	projectsCount := len(projects)
	return &dto.ProjectsResponse{
		Projects: projects[startIndex:endIndex],
		PageInfo: dto.PageInfo{
			PageCount:     int(math.Ceil(float64(projectsCount) / float64(count))),
			CurrentPage:   page,
			ProjectsCount: projectsCount,
		},
	}, nil
}

func (c *Client) FetchIssues(id int, timeToWaitMs int) ([]dto.Issue, error) {
	projectId := strconv.Itoa(id)
	httpClient := &http.Client{}
	response, err := httpClient.Get(c.baseURL +
		"/rest/api/2/search?jql=project=" + projectId + "&expand=changelog&startAt=0&maxResults=1")

	if err != nil || response.StatusCode != http.StatusOK {
		logrus.Error("Unable to get issues for project " + projectId)
		return nil, errorlib.ErrHttpBadGateway
	}

	body, _ := io.ReadAll(response.Body)
	var issueResponse dto.IssuesList
	if err = json.Unmarshal(body, &issueResponse); err != nil {
		logrus.Error("Error while unmarshalling issue response: %w", err)
		return nil, errorlib.ErrHttpInternal
	}

	totalIssuesCount := issueResponse.IssuesCount

	if totalIssuesCount == 0 {
		return make([]dto.Issue, 0), nil
	}

	issues := make([]dto.Issue, 1)
	issues = append(issues, issueResponse.Issues[0])

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
							projectId + "&expand=changelog&startAt=" + strconv.Itoa(startAt) +
							"&maxResults=" + strconv.Itoa(issuesPerRequest)

						response, requestErr := httpClient.Get(requestString)
						body, responseReadErr := io.ReadAll(response.Body)

						if requestErr != nil || responseReadErr != nil {
							wasError = true
							logrus.Error(requestErr, responseReadErr)
							close(stop)
							return
						}

						var issueResponse dto.IssuesList
						err = json.Unmarshal(body, &issueResponse)
						if err != nil {
							logrus.Error(err)
							close(stop)
							return
						}

						mutex.Lock()
						for _, elem := range issueResponse.Issues {
							issues = append(issues, elem)
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
			projectId + "\", waiting now" + strconv.Itoa(timeToWaitMs) + "ms")

		if time.Duration(newTimeToSleep) > c.maxTimeSleep {
			return nil, errorlib.ErrHttpGatewayTimeout
		}
		return c.FetchIssues(id, newTimeToSleep)
	}

	return issues, nil
}
