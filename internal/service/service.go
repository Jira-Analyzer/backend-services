package service

import "github.com/Jira-Analyzer/backend-services/internal/repository"

type Services struct {
	IssueService   *IssueService
	ProjectService *ProjectService
}

func NewServices(repos repository.Repositories) *Services {
	return &Services{
		IssueService:   NewIssueService(repos.IssueRepository),
		ProjectService: NewProjectService(repos.ProjectRepository),
	}
}
