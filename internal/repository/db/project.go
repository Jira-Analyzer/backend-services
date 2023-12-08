package db

import (
	"fmt"

	"github.com/Jira-Analyzer/backend-services/internal/db/models"
	provider "github.com/Jira-Analyzer/backend-services/internal/db/postgres"
)

type ProjectRepository struct {
	db *provider.Provider
}

func NewProjectRepository(provider *provider.Provider) *ProjectRepository {
	return &ProjectRepository{
		db: provider,
	}
}

func (repository *ProjectRepository) GetProjects() ([]models.Project, error) {
	projects := []models.Project{}
	if err := repository.db.Select(&projects, `SELECT * FROM "Project"`); err != nil {
		return nil, fmt.Errorf("Failed to fetch projects due to: %w", err)
	}
	return projects, nil
}
