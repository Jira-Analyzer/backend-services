package service

import (
	"context"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
	"github.com/Jira-Analyzer/backend-services/internal/repository"
)

type ProjectService struct {
	repo repository.IProjectRepository
}

func NewProjectService(repo repository.IProjectRepository) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (service *ProjectService) GetProjects(ctx context.Context) ([]domain.Project, error) {
	return service.repo.GetProjects(ctx)
}

func (service *ProjectService) GetProjectsByRange(ctx context.Context, offset int, count int) ([]domain.Project, error) {
	return service.repo.GetProjectsByRange(ctx, offset, count)
}
