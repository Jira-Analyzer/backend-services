package repository

import (
	"context"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
)

//go:generate mockgen --build_flags=--mod=mod -destination mock/mock_repository.go . IIssueRepository,IProjectRepository

type IIssueRepository interface {
	GetIssuesByProject(context context.Context, projectId int) ([]domain.Issue, error)
}

type IProjectRepository interface {
	GetProjects(context context.Context) ([]domain.Project, error)
}
