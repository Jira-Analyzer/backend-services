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

func (repository *ProjectRepository) GetProjectById(ctx context.Context, id int) (*domain.Project, error) {
	var project domain.Project
	if err := repository.db.GetContext(ctx, &project, `SELECT * FROM "Project" WHERE id=$1`, id); err != nil {
		return nil, fmt.Errorf("failed to fetch projects due to: %w", errorlib.ErrHttpInternal)
	}

	return &project, nil
}

func (repository *ProjectRepository) GetProjectsByRange(ctx context.Context, offset int, count int) ([]domain.Project, error) {
	projects := make([]domain.Project, 0)
	if err := repository.db.SelectContext(ctx, &projects, `SELECT * FROM "Project" ORDER BY id OFFSET $1 ROWS FETCH FIRST $2 ROW ONLY`, offset, count); err != nil {
		return nil, fmt.Errorf("failed to fetch projects due to: %w", err)
	}

	return projects, nil
}

func (repository *ProjectRepository) InsertProject(ctx context.Context, project *domain.Project) error {
	query := `
		INSERT INTO "Project" (id, name, description, avatar_url, type, archived)
		VALUES (:id, :name, :description, :avatar_url, :type, :archived);
	`
	_, err := repository.db.NamedExecContext(ctx, query, project)
	return err
}

func (repository *ProjectRepository) UpdateProject(ctx context.Context, project *domain.Project) error {
	query := `
		UPDATE "Project"
		SET name=:name, description=:description, avatar_url=:avatar_url, type=:type, archived=:archived
		WHERE id=:id;
	`

	_, err := repository.db.NamedExecContext(ctx, query, project)
	return err
}
