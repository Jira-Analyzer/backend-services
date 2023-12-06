package service

import "github.com/Jira-Analyzer/backend-services/internal/repository"

type Services struct {
	IssueService   IIssueService
	ProjectService IProjectService
}

func NewServices(repo *repository.Repositories) *Services {
	return &Services{
		IssueService:   NewIssueService(repo.IssueRepository),
		ProjectService: NewProjectService(repo.ProjectRepository),
	}
}
