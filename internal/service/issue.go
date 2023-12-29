package service

import (
	"github.com/Jira-Analyzer/backend-services/internal/domain"
	"github.com/Jira-Analyzer/backend-services/internal/repository"
)

type IssueService struct {
	repo repository.IssueRepositoryInterface
}

func NewIssueService(repo repository.IssueRepositoryInterface) *IssueService {
	return &IssueService{
		repo: repo,
	}
}

func (service *IssueService) GetIssuesByProject(projectId int) ([]domain.Issue, error) {
	list, err := service.repo.GetIssuesByProject(projectId)
	if err != nil {
		return nil, err
	}

	return list, nil
}
