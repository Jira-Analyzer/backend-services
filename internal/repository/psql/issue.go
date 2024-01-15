package psql

import (
	"context"
	"fmt"

	provider "github.com/Jira-Analyzer/backend-services/internal/db"
	"github.com/Jira-Analyzer/backend-services/internal/domain"
	errorlib "github.com/Jira-Analyzer/backend-services/internal/error"
)

type IssueRepository struct {
	db *provider.Provider
}

func NewIssueRepository(provider *provider.Provider) *IssueRepository {
	return &IssueRepository{
		db: provider,
	}
}

func (repository *IssueRepository) GetIssuesByProject(ctx context.Context, projectId int) ([]domain.Issue, error) {
	issues := make([]domain.Issue, 0)
	if err := repository.db.SelectContext(ctx, &issues, `SELECT * FROM "Issue" WHERE project_id=$1`, projectId); err != nil {
		return nil, fmt.Errorf("Failed to fetch project '%d' issues due to: %w", projectId, errorlib.InternalError)
	}
	return issues, nil
}
