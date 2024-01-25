package service

import (
	"context"
	"time"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
)

//go:generate mockgen --build_flags=--mod=mod -destination mock/mock_service.go . IIssueService,IProjectService

type IIssueService interface {
	GetIssuesByProject(ctx context.Context, projectId int) ([]domain.Issue, error)
	FilterOpenedIssues(ctx context.Context, issues []domain.Issue) []domain.Issue
	FilterClosedIssues(ctx context.Context, issues []domain.Issue) []domain.Issue
	GetAverageTimeSpent(ctx context.Context, issues []domain.Issue) time.Duration
	GetWeekAverageCreatedNumber(ctx context.Context, issues []domain.Issue) int
}

type IProjectService interface {
	GetProjects(ctx context.Context) ([]domain.Project, error)
	GetProjectsByRange(ctx context.Context, offset int, count int) ([]domain.Project, error)
	GetProjectById(ctx context.Context, id int) (*domain.Project, error)
}
