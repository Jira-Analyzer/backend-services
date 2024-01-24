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
		return nil, fmt.Errorf("failed to fetch projects due to: %w", errorlib.ErrHttpInternal)
	}

	return projects, nil
}

func (repository *ProjectRepository) GetProjectsByRange(ctx context.Context, offset int, count int) ([]domain.Project, error) {
	projects := make([]domain.Project, 0)
	if err := repository.db.SelectContext(ctx, &projects, `SELECT * FROM "Project" ORDER BY id OFFSET $1 ROWS FETCH FIRST $2 ROW ONLY`, offset, count); err != nil {
		return nil, fmt.Errorf("failed to fetch projects due to: %w", err)
	}

	return projects, nil
}

func (repository *ProjectRepository) InsertProject(ctx context.Context, project domain.Project) (int, error) {
	query := `
		INSERT INTO projects (name, description, avatar_url, type, archived)
		VALUES (:name, :description, :avatar_url, :type, :archived)
		RETURNING id;
	`

	var insertedID int
	rows, err := repository.db.NamedQueryContext(ctx, query, project)
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

func (repository *ProjectRepository) UpdateProject(ctx context.Context, project domain.Project) error {
	query := `
		UPDATE projects
		SET name=:name, description=:description, avatar_url=:avatar_url, type=:type, archived=:archived
		WHERE id=:id;
	`

	_, err := repository.db.NamedExecContext(ctx, query, project)
	return err
}
