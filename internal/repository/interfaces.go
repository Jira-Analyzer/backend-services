package repository

import "github.com/Jira-Analyzer/backend-services/internal/models"

type IssueRepositoryInterface interface {
	GetIssuesByProject(projectId int) ([]models.Issue, error)
}

type ProjectRepositoryInrerface interface {
	GetProjects() ([]models.Project, error)
}
