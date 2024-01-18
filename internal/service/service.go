package service

import "github.com/Jira-Analyzer/backend-services/internal/repository"

type Services struct {
	IssueService   IIssueService
	ProjectService IProjectService
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		IssueService:   NewIssueService(repos.IssueRepository),
		ProjectService: NewProjectService(repos.ProjectRepository),
	}
}
