package service

import (
	"context"

	"github.com/Jira-Analyzer/backend-services/internal/domain"
	"github.com/Jira-Analyzer/backend-services/internal/repository"
)

type IssueService struct {
	repo repository.IIssueRepository
}

func NewIssueService(repo repository.IIssueRepository) *IssueService {
	return &IssueService{
		repo: repo,
	}
}

func (service *IssueService) GetIssuesByProject(ctx context.Context, projectId int) ([]domain.Issue, error) {
	return service.repo.GetIssuesByProject(ctx, projectId)
}
