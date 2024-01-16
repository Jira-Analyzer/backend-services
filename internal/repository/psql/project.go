package psql

import (
	"context"
	"fmt"

	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	"github.com/Jira-Analyzer/backend-services/internal/domain"
	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
)

type ProjectRepository struct {
	db *provider.Provider
}

func NewProjectRepository(provider *provider.Provider) *ProjectRepository {
	return &ProjectRepository{
		db: provider,
	}
}

func (repository *ProjectRepository) GetProjects(ctx context.Context) ([]domain.Project, error) {
	projects := make([]domain.Project, 0)
	if err := repository.db.SelectContext(ctx, &projects, `SELECT * FROM "Project"`); err != nil {
		return nil, fmt.Errorf("Failed to fetch projects due to: %w", errorlib.InternalError)
	}
	return projects, nil
}
