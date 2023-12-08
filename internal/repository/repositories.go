package repository

import (
	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	db "github.com/Jira-Analyzer/backend-services/internal/repository/psql"
)

type Repositories struct {
	IssueRepositoryInterface
	ProjectRepositoryInrerface
}

func NewRepositories(provider *provider.Provider) *Repositories {
	return &Repositories{
		IssueRepositoryInterface:   db.NewIssueRepository(provider),
		ProjectRepositoryInrerface: db.NewProjectRepository(provider),
	}
}
