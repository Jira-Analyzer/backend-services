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
		return nil, fmt.Errorf("Failed to fetch project '%d' issues due to: %w", projectId, errorlib.ErrHttpInternal)
	}
	return issues, nil
}

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

func (repository *IssueRepository) InsertIssue(ctx context.Context, issue domain.Issue) (int, error) {
	query := `
		INSERT INTO issues (project_id, author_id, reporter_id, key, summary, type, priority, status, created_time, closed_time, updated_time, time_spent)
		VALUES (:project_id, :author_id, :reporter_id, :key, :summary, :type, :priority, :status, :created_time, :closed_time, :updated_time, :time_spent)
		RETURNING id;
	`

	var insertedID int
	rows, err := repository.db.NamedQueryContext(ctx, query, issue)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&insertedID); err != nil {
			return 0, err
		}
	}

	return insertedID, nil
}

func (repository *IssueRepository) UpdateIssue(ctx context.Context, issue domain.Issue) error {
	query := `
		UPDATE issues
		SET project_id=:project_id, author_id=:author_id, reporter_id=:reporter_id, key=:key, summary=:summary, type=:type, priority=:priority, status=:status, created_time=:created_time, closed_time=:closed_time, updated_time=:updated_time, time_spent=:time_spent
		WHERE id=:id;
	`

	_, err := repository.db.NamedExecContext(ctx, query, issue)
	return err
}
