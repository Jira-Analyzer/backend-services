package service

import "github.com/Jira-Analyzer/backend-services/internal/domain"

type IssueServiceInterface interface {
	GetIssuesByProject(projectId int64) ([]domain.Issue, error)
}

type ProjectServiceInrerface interface {
	GetProjects() ([]domain.Project, error)
	GetNFirstProjects(number int) ([]domain.Project, error)
}
