package service

import (
	"context"
	"github.com/Jira-Analyzer/backend-services/internal/domain"
	"github.com/Jira-Analyzer/backend-services/internal/repository"
)

type ProjectService struct {
	projectRepo repository.IProjectRepository
}

func NewProjectService(repo repository.IProjectRepository) *ProjectService {
	return &ProjectService{
		projectRepo: repo,
	}
}

func (service *ProjectService) InsertProject(ctx context.Context, project domain.Project) (int, error) {
	return service.projectRepo.InsertProject(ctx, project)
}

func (service *ProjectService) UpdateProject(ctx context.Context, project domain.Project) error {
	return service.projectRepo.UpdateProject(ctx, project)
}
