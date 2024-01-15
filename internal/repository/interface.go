package repository

import (
	"context"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
)

type IIssueRepository interface {
	GetIssuesByProject(context context.Context, projectId int) ([]domain.Issue, error)
}

type IProjectRepository interface {
	GetProjects(context context.Context) ([]domain.Project, error)
}
