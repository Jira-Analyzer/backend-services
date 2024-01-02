package repository

import (
	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	db "github.com/Jira-Analyzer/backend-services/internal/repository/psql"
)

type Repositories struct {
	IssueRepository   IssueRepositoryInterface
	ProjectRepository ProjectRepositoryInterface
}

func NewRepositories(provider *provider.Provider) *Repositories {
	return &Repositories{
		IssueRepository:   db.NewIssueRepository(provider),
		ProjectRepository: db.NewProjectRepository(provider),
	}
}
