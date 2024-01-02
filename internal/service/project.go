package service

import (
	"fmt"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
	"github.com/Jira-Analyzer/backend-services/internal/repository"
)

type ProjectService struct {
	repo repository.ProjectRepositoryInterface
}

func NewProjectService(repo repository.ProjectRepositoryInterface) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (service *ProjectService) GetProjects() ([]domain.Project, error) {
	list, err := service.repo.GetProjects()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (service *ProjectService) GetNFirstProjects(number int) ([]domain.Project, error) {
	list, err := service.repo.GetProjects()
	if err != nil {
		return nil, err
	}
	length := len(list)
	if number <= 0 || number >= length {
		return nil, fmt.Errorf("expected positive number less than %d (total project count), got: %d", length, number)
	}

	return list[:number], nil
}
