package repository

import (
	"context"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
)

//go:generate mockgen --build_flags=--mod=mod -destination mock/mock_repository.go . IIssueRepository,IProjectRepository

type IIssueRepository interface {
	GetIssuesByProject(context context.Context, projectId int) ([]domain.Issue, error)
	InsertIssue(ctx context.Context, issue domain.Issue) (int, error)
	UpdateIssue(ctx context.Context, issue domain.Issue) error
}

type IProjectRepository interface {
	GetProjects(context context.Context) ([]domain.Project, error)
	GetProjectsByRange(ctx context.Context, offset int, count int) ([]domain.Project, error)
	GetProjectById(ctx context.Context, id int) (*domain.Project, error)
	InsertProject(ctx context.Context, project domain.Project) (int, error)
	UpdateProject(ctx context.Context, project domain.Project) error
}
