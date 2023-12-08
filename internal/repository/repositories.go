package repository

import (
	provider "github.com/Jira-Analyzer/backend-services/internal/db/postgres"
	db "github.com/Jira-Analyzer/backend-services/internal/repository/db"
)

type Repositories struct {
	*db.IssueRepository
	*db.ProjectRepository
}

func NewRepositories(provider *provider.Provider) *Repositories {
	return &Repositories{
		IssueRepository:   db.NewIssueRepository(provider),
		ProjectRepository: db.NewProjectRepository(provider),
	}
}
