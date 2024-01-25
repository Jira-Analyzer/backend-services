package service

import (
	"context"
	"time"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
	"github.com/Jira-Analyzer/backend-services/internal/repository"
	"github.com/ledongthuc/goterators"
)

type IssueService struct {
	repo repository.IIssueRepository
}

func NewIssueService(repo repository.IIssueRepository) *IssueService {
	return &IssueService{
		repo: repo,
	}
}

func (service *IssueService) GetIssuesByProject(ctx context.Context, projectId int) ([]domain.Issue, error) {
	return service.repo.GetIssuesByProject(ctx, projectId)
}

func (service *IssueService) FilterOpenedIssues(ctx context.Context, issues []domain.Issue) []domain.Issue {
	return goterators.Filter(issues, func(item domain.Issue) bool { return item.Status == "Opened" })
}

func (service *IssueService) FilterClosedIssues(ctx context.Context, issues []domain.Issue) []domain.Issue {
	return goterators.Filter(issues, func(item domain.Issue) bool { return item.Status == "Closed" })
}

func (service *IssueService) GetAverageTimeSpent(ctx context.Context, issues []domain.Issue) time.Duration {
	if len(issues) == 0 {
		return time.Duration(0)
	}

	var totalTime time.Duration
	for _, issue := range issues {
		totalTime += issue.TimeSpent
	}
	return time.Duration(totalTime.Nanoseconds() / int64(len(issues)))
}

func (service *IssueService) GetWeekAverageCreatedNumber(ctx context.Context, issues []domain.Issue) int {
	var totalNumber int
	current := time.Now()
	begin := time.Now().AddDate(0, 0, -7)
	for _, issue := range issues {
		if issue.CreatedTime.After(begin) && issue.CreatedTime.Before(current) {
			totalNumber++
		}
	}
	return totalNumber / 7
}
