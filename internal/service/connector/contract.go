package service

import (
	"context"
	"github.com/Jira-Analyzer/backend-services/internal/domain"
)

//go:generate mockgen --build_flags=--mod=mod -destination mock/mock_service.go . IIssueService,IProjectService

type IIssueService interface {
	InsertIssue(ctx context.Context, issue domain.Issue) (int, error)
	UpdateIssue(ctx context.Context, issue domain.Issue) error
}

type IProjectService interface {
	InsertProject(ctx context.Context, project domain.Project) (int, error)
	UpdateProject(ctx context.Context, project domain.Project) error
}
