package psql

import (
	"fmt"

	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	"github.com/Jira-Analyzer/backend-services/internal/models"
)

type IssueRepository struct {
	db *provider.Provider
}

func NewIssueRepository(provider *provider.Provider) *IssueRepository {
	return &IssueRepository{
		db: provider,
	}
}

func (repository *IssueRepository) GetIssuesByProject(projectId int) ([]models.Issue, error) {
	issues := []models.Issue{}
	if err := repository.db.Select(&issues, `SELECT * FROM "Issue" WHERE project_id=$1`, projectId); err != nil {
		return nil, fmt.Errorf("Failed to fetch project '%d' issues due to: %w", projectId, err)
	}
	return issues, nil
}
