package service

import (
	"context"
	"github.com/Jira-Analyzer/backend-services/internal/domain"
	"github.com/Jira-Analyzer/backend-services/internal/repository"
)

type IssueService struct {
	issueRepo repository.IIssueRepository
}

func NewIssueService(issueRepo repository.IIssueRepository) *IssueService {
	return &IssueService{
		issueRepo: issueRepo,
	}
}

func (service *IssueService) InsertIssue(ctx context.Context, issue domain.Issue) (int, error) {
	return service.issueRepo.InsertIssue(ctx, issue)
}

func (service *IssueService) UpdateIssue(ctx context.Context, issue domain.Issue) error {
	return service.issueRepo.UpdateIssue(ctx, issue)
}
