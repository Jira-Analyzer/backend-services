package dto

import (
	"strconv"
	"time"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
)

type IssuesList struct {
	IssuesCount int     `json:"total"`
	Issues      []Issue `json:"issues"`
}

type Issue struct {
	Id     string      `json:"id"`
	Key    string      `json:"key"`
	Fields IssueFields `json:"fields"`
}

type IssueFields struct {
	Summary string `json:"summary"`
	Type    struct {
		Name string `json:"name"`
	} `json:"issuetype"`
	Status struct {
		Name string `json:"name"`
	} `json:"status"`
	Priority struct {
		Name string `json:"name"`
	} `json:"priority"`
	Creator struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
	} `json:"creator"`
	Reporter struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
	} `json:"reporter"`
	Project struct {
		Name string `json:"name"`
	} `json:"project"`
	Description  string `json:"description"`
	AssigneeName struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
	} `json:"assignee"`
	CreatedTime string `json:"created"`
	UpdatedTime string `json:"updated"`
	ClosedTime  string `json:"resolutiondate"`
}

type IssueStatusChange struct {
	Changelog struct {
		Histories []History `json:"histories"`
	} `json:"changelog"`
}

type History struct {
	Author struct {
		Name string `json:"name"`
	} `json:"author"`
	CreatedTime string `json:"created"`
	Items       []Item
}

type Item struct {
	Field      string `json:"field"`
	FromString string `json:"fromString"`
	ToString   string `json:"toString"`
}

func (issue *Issue) ToDomainIssue(projectId int) domain.Issue {
	createdTime, _ := time.Parse("2006-01-02T15:04:05.999-0700", issue.Fields.CreatedTime)
	updatedTime, _ := time.Parse("2006-01-02T15:04:05.999-0700", issue.Fields.UpdatedTime)
	closedTime, _ := time.Parse("2006-01-02T15:04:05.999-0700", issue.Fields.ClosedTime)
	timeSpent := time.Hour * 0
	if len(issue.Fields.ClosedTime) == 0 {
		closedTime = createdTime.Add(time.Hour * 24 * 365 * 10)
	} else {
		timeSpent = time.Duration(closedTime.Sub(createdTime).Hours())
	}
	id, _ := strconv.Atoi(issue.Id)
	return domain.Issue{
		Id:          id,
		ProjectId:   projectId,
		Key:         issue.Key,
		Summary:     issue.Fields.Summary,
		Priority:    issue.Fields.Priority.Name,
		Status:      issue.Fields.Status.Name,
		ClosedTime:  closedTime,
		CreatedTime: createdTime,
		UpdatedTime: updatedTime,
		TimeSpent:   timeSpent,
		Author:      issue.Fields.Creator.Name,
		Reporter:    issue.Fields.Reporter.Name,
	}
}
