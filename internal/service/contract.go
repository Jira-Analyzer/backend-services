package service

import (
	"context"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
)

//go:generate mockgen --build_flags=--mod=mod -destination mock/mock_service.go . IIssueService,IProjectService

type IIssueService interface {
	GetIssuesByProject(ctx context.Context, projectId int64) ([]domain.Issue, error)
}

type IProjectService interface {
	GetProjects(ctx context.Context) ([]domain.Project, error)
	GetProjectsByRange(ctx context.Context, offset int, count int) ([]domain.Project, error)
}
