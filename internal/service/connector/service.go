package service

import (
	"github.com/Jira-Analyzer/backend-services/internal/client/jira"
	"github.com/Jira-Analyzer/backend-services/internal/repository"
)

type Service struct {
	client       *jira.Client
	issuesRepo   repository.IIssueRepository
	projectsRepo repository.IProjectRepository
}

func NewService(config *jira.Config, repositories *repository.Repositories) *Service {
	return &Service{
		client:       jira.NewClient(config),
		issuesRepo:   repositories.IssueRepository,
		projectsRepo: repositories.ProjectRepository,
	}
}
