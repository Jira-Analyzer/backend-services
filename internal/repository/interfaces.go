package repository

import "github.com/Jira-Analyzer/backend-services/internal/domain"

type IssueRepositoryInterface interface {
	GetIssuesByProject(projectId int) ([]domain.Issue, error)
}

type ProjectRepositoryInterface interface {
	GetProjects() ([]domain.Project, error)
}
