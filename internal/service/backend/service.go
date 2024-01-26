package service

import (
	"github.com/Jira-Analyzer/backend-services/internal/repository"
	server "github.com/Jira-Analyzer/backend-services/internal/server/backend"
)

type Services struct {
	IssueService   IIssueService
	ProjectService IProjectService
	Proxy          *ProxyService
}

func NewServices(repos *repository.Repositories, conf *server.ServerConfig) *Services {
	return &Services{
		IssueService:   NewIssueService(repos.IssueRepository),
		ProjectService: NewProjectService(repos.ProjectRepository),
		Proxy:          NewProxyService(conf),
	}
}
